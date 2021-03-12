package dhtnode

import (
	"context"
	"errors"
	"log"
	"time"

	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/zebra-uestc/chord"
	bm "github.com/zebra-uestc/chord/models/bridge"
	cm "github.com/zebra-uestc/chord/models/chord"
	"google.golang.org/grpc"
	"github.com/zebra-uestc/chord/config"
)

type dhtNode struct {
	IsMainNode            bool
	exitChan              chan struct{}
	pendingBatch          []*cb.Envelope
	pendingBatchSizeBytes uint32
	PendingBatchStartTime time.Time
	ChannelID             string
	*chord.Node
	mn mainNodeInside
	// Metrics   *Metrics
}

// NewDhtNode 创建DhtNode
func NewDhtNode(cnf *chord.Config, joinNode *cm.Node) (*dhtNode, error) {
	node, err := chord.NewNode(cnf, joinNode)
	dhtnode := &dhtNode{Node: node}

	if err != nil {
		log.Println("transport start error:", err)
		return nil, err
	}

	txStore, ok := dhtnode.GetStorage().(chord.TxStorage)
	if !ok {
		log.Fatal("Storage Error")
		return nil, errors.New("Storage Error")
	}
	sendMsgChan := txStore.GetMsgChan()

	//生成prevBlock
	go dhtnode.PrevBlock(sendMsgChan)
	return dhtnode, err
}

func (dhtn *dhtNode) DhtInsideTransBlock(block *cb.Block) error {
	if !dhtn.IsMainNode {
		conn, err := grpc.Dial(config.MainNodeAddressMsg, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c := bm.NewBlockTranserClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
		defer cancel()

		finalBlockByte, err := protoutil.Marshal(block)
		if err != nil {
			log.Fatalf("marshal err")
		}
		_, err = c.TransBlock(ctx, &bm.BlockBytes{BlockPayload: finalBlockByte})
		if err != nil {
			log.Fatalf("could not transcation Block: %v", err)
		}
		conn.Close()
		return err
	}
	// 将生成的Block放到mainNode底下的Channel中
	dhtn.mn.SendPrevBlockToChan(block)
	return nil
}


// PrevBlock 将区块进行排序并发送给orderer
func (dhtn *dhtNode) PrevBlock(sendMsgChan chan *chord.Message) {
	var timer <-chan time.Time

	var cnt uint = 0 // test

	for {
		select {
		case msg, ok := <-sendMsgChan:
			if !ok {
				println("channel sendMsgChan is closed!")
			}
			if msg.ConfigMsg == nil || msg.ConfigMsg.Payload == nil {
				batches, pending := dhtn.Ordered(msg.NormalMsg)
				//出块并发送给mainnode或者orderer
				for _, batch := range batches {
					block := dhtn.PreCreateNextBlock(batch)

					cnt = cnt + 1
					// println("PreBlock", cnt)

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

		case <-dhtn.GetShutdownCh():
			logger.Debugf("Exiting")
			return
		default:
			// do nothing
		}
	}
}
