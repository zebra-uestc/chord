package dhtnode

import (
	"context"
	"crypto/sha256"

	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/zebra-uestc/chord"
	bm "github.com/zebra-uestc/chord/models/bridge"
	"google.golang.org/grpc"
)

var (
	eMptyRequest       = &bm.DhtStatus{}
	OrdererAddress     = ":6666"
	GenesisesBlockHash = []byte{0x09, 0xdd, 0xec, 0x54, 0x73, 0xcd, 0xaa, 0x05, 0x39, 0xe8, 0x37, 0x75, 0x48, 0x32, 0x35, 0x6d, 0x13, 0x34, 0xf1, 0xde, 0x89, 0x26, 0x18, 0xaa, 0x42, 0x2b, 0x9b, 0x5b, 0xfb, 0xd0, 0x0f, 0x77}
)

type MainNode interface {
	AddNode(id string, addr string) error
	StartDht(id string, address string)
	StartTransMsgServer(address string)
	StartTransBlockServer(address string)
	Stop()
}

type server struct{}

type mainNode struct {
	*dhtNode
	prevBlockChan chan *cb.Block
	sendBlockChan chan *cb.Block
	bm.UnimplementedBlockTranserServer
	bm.UnimplementedMsgTranserServer
	//lastBlockCnf *bm.Config
	lastBlockHash []byte
	blockNum      uint64
}

func NewMainNode() (MainNode, error) {

	//TODO:暂时只定义了接口LoadConfig，还未实现其内容，无法容忍mainnode宕机
	////向orderer询问lastBlock
	//conn, err := grpc.Dial(OrdererAddress, grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//c := bm.NewBlockTranserClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.LoadConfig(ctx, nil)
	//mainNode.lastBlockCnf = r
	//if err != nil {
	//	log.Fatalf("could not transcation Block: %v", err)
	//}

	return &mainNode{
		prevBlockChan: make(chan *cb.Block, 10),
		sendBlockChan: make(chan *cb.Block, 10),
		blockNum:      1,
	}, nil
}

type mainNodeInside interface {
	SendPrevBlockToChan(*cb.Block)
}

func (mn *mainNode) StartDht(id, address string) {
	nodeCnf := chord.DefaultConfig()
	nodeCnf.Id = id
	nodeCnf.Addr = address
	nodeCnf.Timeout = 10 * time.Millisecond
	nodeCnf.MaxIdle = 100 * time.Millisecond
	node, _ := NewDhtNode(nodeCnf, nil)
	node.mn = mn // main_node继承dht_node，要给dht_node里面的mn变量初始化
	mn.dhtNode = node
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
	//给orderer发Block
	go func() {
		ticker := time.NewTicker(1 * time.Second)

		for {
			// todo
			select {
			//给区块编号
			case prevBlock := <-mn.prevBlockChan:
				println("get block")
				newBlock := mn.FinalBlock(prevBlock)
				//将新生成的块放到sendBlockChan转发给orderer
				mn.sendBlockChan <- newBlock
				//更新最后一个区块的哈希和区块个数
				mn.lastBlockHash = protoutil.BlockHeaderHash(newBlock.Header)
				mn.blockNum++

			case finalBlock := <-mn.sendBlockChan:

				conn, err := grpc.Dial(OrdererAddress, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				c := bm.NewBlockTranserClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
				defer cancel()
				finalBlockByte, err := protoutil.Marshal(finalBlock)
				if err != nil {
					log.Fatalf("marshal err")
				}
				_, err = c.TransBlock(ctx, &bm.BlockBytes{BlockPayload: finalBlockByte})

			case <-mn.GetShutdownCh():
				ticker.Stop()
			}
		}
	}()

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

func (mn *mainNode) AddNode(id string, addr string) error {
	cnf := chord.DefaultConfig()
	cnf.Id = id
	cnf.Addr = addr
	cnf.Timeout = 10 * time.Millisecond
	cnf.MaxIdle = 100 * time.Millisecond
	_, err := NewDhtNode(cnf, mn.dhtNode.Node.Node)
	return err
}

// order To dht的处理
func (mn *mainNode) TransMsg(ctx context.Context, msg *bm.MsgBytes) (*bm.DhtStatus, error) {

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

//接收其他节点的block
func (mainNode *mainNode) TransBlock(ctx context.Context, blockByte *bm.BlockBytes) (*bm.DhtStatus, error) {
	//反序列化为Block
	// block := &cb.Block{}
	block, err := protoutil.UnmarshalBlock(blockByte.BlockPayload)
	if err != nil {
		return nil, err
	}
	mainNode.prevBlockChan <- block
	return &bm.DhtStatus{}, nil
}

//给区块编号
func (mainNode *mainNode) FinalBlock(block *cb.Block) *cb.Block {
	block.Header.PreviousHash = mainNode.lastBlockHash
	if mainNode.blockNum == 1 {
		block.Header.PreviousHash = GenesisesBlockHash
	}
	block.Header.Number = mainNode.blockNum
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
	close(mn.GetShutdownCh())
}
