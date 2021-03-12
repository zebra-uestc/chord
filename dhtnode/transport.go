package dhtnode

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	bm "github.com/zebra-uestc/chord/models/bridge"

	"google.golang.org/grpc"
)

// 与chord原本的transport.go不同的是：我们只实现client发起连接的持久化，而暂时不管server的监听，因为监听没有可优化的地方

func Dial(addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, opts...)
}

// Client(outbound) dial config
type Config struct {
	// Id   string
	// Addr string

	DialOpts []grpc.DialOption

	Timeout time.Duration // 作为WithTimeout()函数参数
	MaxIdle time.Duration // 超过maxidle自动关闭连接
}

func DefaultConfig() *Config {
	n := &Config{
		DialOpts: make([]grpc.DialOption, 0, 5),
		Timeout:  10 * time.Second,
		MaxIdle:  10,
	}
	n.DialOpts = append(n.DialOpts,
		grpc.WithBlock(),
		//修改超时时间
		grpc.WithTimeout(1*time.Second), // 与config.Timeout不一样
		grpc.FailOnNonTempDialError(true),
		grpc.WithInsecure(),
	)
	return n
}

type GrpcTransport struct {
	config *Config // 根据config中的id来创建server
	// *bm.UnimplementedBlockTranserServer
	timeout time.Duration
	maxIdle time.Duration

	pool    map[string]*grpcConn
	poolMtx sync.RWMutex

	// server *grpc.Server

	shutdown int32 // 关闭所有连接的标志位
}

// func NewGrpcTransport(config *Config) (cm.ChordClient, error) {
func NewGrpcTransport() *GrpcTransport {

	// addr := config.Addr
	// Try to start the listener
	// listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	return nil, err
	// }
	config := DefaultConfig()

	pool := make(map[string]*grpcConn)

	// Setup the transport
	grp := &GrpcTransport{
		// sock:    listener.(*net.TCPListener),
		timeout: config.Timeout,
		maxIdle: config.MaxIdle,
		pool:    pool,
		config:  config,
	}

	// grp.server = grpc.NewServer(config.ServerOpts...)

	// Done
	return grp
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
	conn, err = Dial(addr, g.config.DialOpts...)
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

// Returns an outbound TCP connection to the pool
func (g *GrpcTransport) returnConn(o *grpcConn) {
	// Update the last asctive time
	o.lastActive = time.Now()

	// Push back into the pool
	g.poolMtx.Lock()
	defer g.poolMtx.Unlock()
	if atomic.LoadInt32(&g.shutdown) == 1 {
		o.conn.Close()
		return
	}
	g.pool[o.addr] = o
}

func (g *GrpcTransport) Start() error {
	// Start RPC server
	// go g.listen()

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
		if time.Since(conn.lastActive) > g.maxIdle {
			conn.Close()
			delete(g.pool, host)
		}
	}
}
