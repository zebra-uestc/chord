package dhtnode

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zebra-uestc/chord/config"
	bm "github.com/zebra-uestc/chord/models/bridge"

	"google.golang.org/grpc"
)

// 与chord原本的transport.go不同的是：我们只实现client发起连接的持久化，而暂时不管server的监听，因为监听没有可优化的地方

// Client(outbound) dial config
type Config struct {
	DialOpts []grpc.DialOption
	Timeout time.Duration // 作为WithTimeout()函数参数
	MaxIdle time.Duration // 超过maxidle自动关闭连接
}

type GrpcTransport struct {
	config *Config // 根据config中的id来创建server
	pool    map[string]*grpcConn
	poolMtx sync.RWMutex
	shutdown int32 // 关闭所有连接的标志位
}

// func NewGrpcTransport(config *Config) (cm.ChordClient, error) {
func NewGrpcTransport() *GrpcTransport {
	return &GrpcTransport{
		pool:    make(map[string]*grpcConn),
	    config:  &Config{
			DialOpts: []grpc.DialOption{
				grpc.WithBlock(),
				grpc.FailOnNonTempDialError(true),
				grpc.WithInsecure(),
			},
			Timeout:  config.GrpcTimeout,
			MaxIdle:  100 * config.GrpcTimeout,
		},
    }
}

type grpcConn struct {
	addr       string
	client     bm.BlockTranserClient
	conn       *grpc.ClientConn
	lastActive time.Time
}

func (g *grpcConn) Close() {
	g.conn.Close()
}

// Gets an outbound connection to a host
func (g *GrpcTransport) getConn(
	addr string,
) (bm.BlockTranserClient, error) {

	g.poolMtx.RLock()
	if atomic.LoadInt32(&g.shutdown) == 1 {
		g.poolMtx.Unlock()
		return nil, fmt.Errorf("TCP transport is shutdown")
	}
	cc, ok := g.pool[addr]
	g.poolMtx.RUnlock()
	
	if ok {
		return cc.client, nil
	}

	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial(addr, g.config.DialOpts...)
	if err != nil {
		return nil, err
	}

	client := bm.NewBlockTranserClient(conn)
	cc = &grpcConn{addr, client, conn, time.Now()}
	
	g.poolMtx.Lock()
	if g.pool == nil {
		g.poolMtx.Unlock()
		return nil, errors.New("must instantiate node before using")
	}
	g.pool[addr] = cc
	g.poolMtx.Unlock()

	return client, nil
}

func (g *GrpcTransport) Start() error {
	// Reap old connections
	go g.reapOld()
	return nil
}

// Close all outbound connection in the pool
func (g *GrpcTransport) Stop() error {
	atomic.StoreInt32(&g.shutdown, 1)

	// Close all the connections
	g.poolMtx.Lock()

	// g.server.Stop()
	for _, conn := range g.pool {
		conn.Close()
	}
	g.pool = nil

	g.poolMtx.Unlock()

	return nil
}

// Closes old outbound connections
func (g *GrpcTransport) reapOld() {
	ticker := time.NewTicker(60 * time.Second)

	for {
		if atomic.LoadInt32(&g.shutdown) == 1 {
			return
		}
		select {
		case <-ticker.C:
			g.reap()
		}

	}
}

func (g *GrpcTransport) reap() {
	g.poolMtx.Lock()
	defer g.poolMtx.Unlock()
	for host, conn := range g.pool {
		if time.Since(conn.lastActive) > g.config.MaxIdle {
			conn.Close()
			delete(g.pool, host)
		}
	}
}
