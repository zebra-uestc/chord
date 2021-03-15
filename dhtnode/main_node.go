package dhtnode

import (
	"context"
	"crypto/sha256"
	"sync"

	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/zebra-uestc/chord"
	"github.com/zebra-uestc/chord/config"
	bm "github.com/zebra-uestc/chord/models/bridge"
	"google.golang.org/grpc"
)

// MainNode 主节点，负责接受Orderer的Msg，通过node的内部机制转发给其它DhtNode
// 接受其它DhtNode及自身的出块，排序后发送给Orderer
type MainNode interface {
	StartDht(id string, address string)
	StartTransMsgServer(address string)
	StartTransBlockServer(address string)
	Stop()
}

type server struct{}

type mainNode struct {
	*DhtNode
	prevBlockChan chan *cb.Block
	sendBlockChan chan *cb.Block
	bm.UnimplementedBlockTranserServer
	bm.UnimplementedMsgTranserServer
	//lastBlockCnf *bm.Config
	lastBlockHash []byte
	blockNum      uint64
	mutex         sync.RWMutex
	isFirstMsg	  bool
	

	Transport *GrpcTransport
}

// NewMainNode 创建mainNode节点
// 其中会向Orderer询问配置，如果无法通信，则等待1分钟
func NewMainNode() (MainNode, error) {
	//向orderer询问lastBlock
	log.Println("Loading Config...")
	conn, err := grpc.Dial(config.OrdererAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := bm.NewBlockTranserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), config.GrpcTimeout)
	defer cancel()
	cnf, err := c.LoadConfig(ctx, &bm.DhtStatus{})
	if err != nil {
		log.Fatalf("could not transcation Block: %v", err)
	}

	conn.Close()

	mn := &mainNode{
		prevBlockChan: make(chan *cb.Block, 10),
		sendBlockChan: make(chan *cb.Block, 1),
		blockNum:      cnf.LastBlockNum,
		lastBlockHash: cnf.PrevBlockHash,
		Transport:     NewGrpcTransport(),
		isFirstMsg:    true,
	}

	//给preblock标号，并给orderer发Block
	go mn.Process()
	// 开启回收旧连接
	go mn.Transport.Start()

	return mn, nil
}

type mainNodeInside interface {
	SendPrevBlockToChan(*cb.Block)
}

func (mn *mainNode) StartDht(id, address string) {
	nodeCnf := chord.DefaultConfig()
	nodeCnf.Id = id
	nodeCnf.Addr = address
	nodeCnf.Timeout = config.GrpcTimeout
	nodeCnf.MaxIdle = 100 * config.GrpcTimeout
	node, _ := NewDhtNode(nodeCnf, nil)
	node.mn = mn // main_node继承dht_node，要给dht_node里面的mn变量初始化
	mn.DhtNode = node
	mn.IsMainNode = true
}

func (mn *mainNode) startTransMsgServer(address string) {
	println("MsgTranserServer listen:", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	s := grpc.NewServer()
	bm.RegisterMsgTranserServer(s, mn)
	println("MsgTranserServer start serve")
	if err := s.Serve(lis); err != nil {
		log.Fatal("fail to  serve:", err)
	}
	println("MsgTranserServer serve end")
}

func (mn *mainNode) StartTransMsgServer(address string) {
	go mn.startTransMsgServer(address)
}

func (mn *mainNode) startTransBlockServer(address string) {

	println("TransBlockServer listen:", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	s := grpc.NewServer()
	bm.RegisterBlockTranserServer(s, mn)
	println("TransBlockServer start serve")
	if err := s.Serve(lis); err != nil {
		log.Fatal("fail to  serve:", err)
	}

	println("TransBlockServer serve end")
}

func (mn *mainNode) StartTransBlockServer(address string) {
	go mn.startTransBlockServer(address)
}

func (mn *mainNode) SendPrevBlockToChan(block *cb.Block) {
	mn.prevBlockChan <- block
}

// order To dht的处理
func (mn *mainNode) TransMsg(ctx context.Context, msg *bm.MsgBytes) (*bm.DhtStatus, error) {
	//记录收到第一条消息的时间
	if msg.NormalMsg != nil {
		if mn.isFirstMsg {
			startTime := time.Now().UnixNano()/1e6
			println(startTime)
			mn.isFirstMsg = false
		}
	}
	// println("get msg")
	value, err := proto.Marshal(msg)

	if err != nil {
		log.Println("Marshal err: ", err)
	}
	//将value加密为hashKey
	hashKey, err := mn.hashValue(value)
	if err != nil {
		log.Println("hashVal err: ", err)
	}
	//通过dht环转发到其他节点并存储在storage里面,并且放在同到Msgchan
	err = mn.Set(hashKey, value)
	return &bm.DhtStatus{}, nil
}

// TransBlock 接收其他节点的block
func (mn *mainNode) TransBlock(ctx context.Context, blockByte *bm.BlockBytes) (*bm.DhtStatus, error) {
	//反序列化为Block
	// block := &cb.Block{}
	block, err := protoutil.UnmarshalBlock(blockByte.BlockPayload)
	if err != nil {
		return nil, err
	}
	mn.prevBlockChan <- block
	return &bm.DhtStatus{}, nil
}

// FinalBlock 给区块编号
func (mn *mainNode) FinalBlock(block *cb.Block) *cb.Block {
	block.Header.PreviousHash = mn.lastBlockHash
	block.Header.Number = mn.blockNum
	return block
}

func (mn *mainNode) hashValue(key []byte) ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write(key); err != nil {
		return nil, err
	}
	hashVal := h.Sum(nil)
	return hashVal, nil
}

func (mn *mainNode) Stop() {
	mn.Node.Stop()
}

func (mn *mainNode) Process() {
	// ticker := time.NewTicker(1 * time.Second)

	for {
		// todo
		select {
		//给区块编号
		case prevBlock, ok := <-mn.prevBlockChan:
			if !ok {
				println("channel prevBlockChan is closed!")
			}
			newBlock := mn.FinalBlock(prevBlock)

			//将新生成的块放到sendBlockChan转发给orderer
			mn.sendBlockChan <- newBlock
			//更新最后一个区块的哈希和区块个数
			mn.mutex.Lock()

			mn.lastBlockHash = protoutil.BlockHeaderHash(newBlock.Header)
			mn.blockNum++
			mn.mutex.Unlock()

			// println("full block", newBlock.Header.Number)

		case finalBlock, ok := <-mn.sendBlockChan:
			if !ok {
				println("channel sendBlockChan is closed!")
			}

			// conn, err := grpc.Dial(config.OrdererAddress, grpc.WithInsecure(), grpc.WithBlock())
			// if err != nil {
			// 	log.Fatalf("did not connect: %v", err)
			// }
			// c := bm.NewBlockTranserClient(conn)
			c, err := mn.Transport.getConn(config.OrdererAddress)

			ctx, cancel := context.WithTimeout(context.Background(), mn.Transport.config.Timeout)
			defer cancel()
			finalBlockByte, err := protoutil.Marshal(finalBlock)
			if err != nil {
				log.Fatalf("marshal err")
			}

			// println("to send4", finalBlock.Header.Number)

			_, err = c.TransBlock(ctx, &bm.BlockBytes{BlockPayload: finalBlockByte})

			//记录每个区块的发送时间
			println("send block time:", time.Now().UnixNano()/1e6)

			// conn.Close()

		// case <-mn.GetShutdownCh():
		// 	ticker.Stop()
		default:
			// do nothing
		}
	}
}
