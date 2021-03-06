package dhtnode

import (
	"context"
	"errors"
	"log"
	"time"

	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/zebra-uestc/chord"
	"github.com/zebra-uestc/chord/config"
	bm "github.com/zebra-uestc/chord/models/bridge"
	cm "github.com/zebra-uestc/chord/models/chord"
)

type DhtNode struct {
	sendBlockChan         chan *cb.Block
	IsMainNode            bool
	pendingBatch          []*cb.Envelope
	pendingBatchSizeBytes uint32
	PendingBatchStartTime time.Time
	ChannelID             string
	*chord.Node
	mn        mainNodeInside
	Transport *GrpcTransport
}

// NewDhtNode 创建DhtNode
func NewDhtNode(cnf *chord.Config, joinNode *cm.Node) (*DhtNode, error) {
	node, err := chord.NewNode(cnf, joinNode)
	if err != nil {
		log.Println("transport start error:", err)
		return nil, err
	}
	dhtnode := &DhtNode{Node: node, Transport: NewGrpcTransport(), sendBlockChan: make(chan *cb.Block, 10)}
	txStore, ok := dhtnode.GetStorage().(chord.TxStorage)
	if !ok {
		log.Fatal("Storage Error")
		return nil, errors.New("Storage Error")
	}
	sendMsgChan := txStore.GetMsgChan()

	//生成prevBlock
	go dhtnode.PrevBlock(sendMsgChan)
	//回收旧连接
	go dhtnode.Transport.Start()
	return dhtnode, err
}

func (dhtn *DhtNode) DhtInsideTransBlock(block *cb.Block) error {
	if !dhtn.IsMainNode {
		dhtn.sendBlockChan <- block
		return nil
	}
	// 将生成的Block放到mainNode底下的Channel中
	dhtn.mn.SendPrevBlockToChan(block)
	return nil
}

func (dhtn *DhtNode) TransPrevBlockClient() error {
	c, err := dhtn.Transport.getConn(config.MainNodeAddressBlock)
	if err != nil {
		log.Fatalln("Can't get conn with main_node: ", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sender, err := c.TransPrevBlock(ctx)
	if err != nil {
		return err
	}

	for preblock := range dhtn.sendBlockChan {

		preBlockByte, err := protoutil.Marshal(preblock)
		if err != nil {
			log.Fatalf("marshal err")
		}
		err = sender.Send(&bm.BlockBytes{BlockPayload: preBlockByte})
		if err != nil {
			log.Fatalln("Can't trans block to orderer:", err)
		}
		// //记录每个区块的发送时间
		// println("send block time:", time.Now().UnixNano()/1e6)

	}
	_, err = sender.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not transcation MsgBytes: %v", err)
	}

	return err
}

// PrevBlock 将区块进行排序并发送给orderer
func (dhtn *DhtNode) PrevBlock(sendMsgChan chan *chord.Message) {
	var timer <-chan time.Time

	var cnt uint = 0 // test
	count := 0

	for {
		select {
		case msg, ok := <-sendMsgChan:
			count = count + 1
			if count == cap(dhtn.sendBlockChan) {
				go dhtn.TransPrevBlockClient()
			}
			if !ok {
				println("channel sendMsgChan is closed!")
			}
			cnt++
			// println("msg count :",cnt)
			if msg.ConfigMsg == nil || msg.ConfigMsg.Payload == nil {
				batches, pending := dhtn.Ordered(msg.NormalMsg)
				//出块并发送给mainnode或者orderer
				for _, batch := range batches {
					block := dhtn.PreCreateNextBlock(batch)

					cnt = cnt + 1
					println("PreBlock", cnt)

					//将PreCreateNextBlock传给MainNode
					err := dhtn.DhtInsideTransBlock(block)
					if err != nil {
						log.Fatalf("could not transcation Block: %v", err)
					}
				}

				switch {
				case timer != nil && !pending:
					// Timer is already running but there are no messages pending, stop the timer
					timer = nil
				case timer == nil && pending:
					// Timer is not already running and there are messages pending, so start it
					//默认时间1s
					timer = time.After(config.BathchTimeout)
					logger.Debugf("Just began %s batch timer", config.BathchTimeout.String())
				default:
					// Do nothing when:
					// 1. Timer is already running and there are messages pending
					// 2. Timer is not set and there are no messages pending
				}
			} else {
				batch := dhtn.Cut()
				if batch != nil {
					block := dhtn.PreCreateNextBlock(batch)
					err := dhtn.DhtInsideTransBlock(&cb.Block{Header: block.Header, Data: block.Data, Metadata: block.Metadata /*参数v*/})
					if err != nil {
						log.Fatalf("could not transcation Block: %v", err)
					}
				}
				block := dhtn.PreCreateNextBlock([]*cb.Envelope{msg.ConfigMsg})
				err := dhtn.DhtInsideTransBlock(&cb.Block{Header: block.Header, Data: block.Data, Metadata: block.Metadata /*参数v*/})
				if err != nil {
					log.Fatalf("could not transcation Block: %v", err)
				}
				timer = nil
			}
		case <-timer:
			//clear the timer
			timer = nil
			batch := dhtn.Cut()
			if len(batch) == 0 {
				logger.Warningf("Batch timer expired with no pending requests, this might indicate a bug")
				continue
			}
			logger.Debugf("Batch timer expired, creating block")
			block := dhtn.PreCreateNextBlock(batch)
			err := dhtn.DhtInsideTransBlock(&cb.Block{Header: block.Header, Data: block.Data, Metadata: block.Metadata /*参数v*/})
			if err != nil {
				log.Fatalf("could not transcation Block: %v", err)
			}
			cnt++
			println("PreBlock", cnt)

		case <-dhtn.GetShutdownCh():
			logger.Debugf("Exiting")
			return
		default:
			// do nothing
		}
	}
}
