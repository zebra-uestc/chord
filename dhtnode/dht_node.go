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
)

var (
	emptyPrevHash  = []byte{}
	toMainNodeAddr = "127.0.0.1:8002"
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

func (dhtn *dhtNode) DhtInsideTransBlock(block *cb.Block) error {

	if !dhtn.IsMainNode {
		conn, err := grpc.Dial(MAINNODEADDRESS, grpc.WithInsecure(), grpc.WithBlock())
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
		return err
	} else {
		// 将生成的Block放到mainNode底下的Channel中
		dhtn.mn.SendPrevBlockToChan(block)

	}
	return nil
}

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

func (dhtnode *dhtNode) PrevBlock(sendMsgChan chan *chord.Message) {
	var timer <-chan time.Time
	for {
		select {
		case msg := <-sendMsgChan:
			if msg.ConfigMsg == nil || msg.ConfigMsg.Payload == nil {
				batches, pending := dhtnode.Ordered(msg.NormalMsg)
				//出块并发送给mainnode或者orderer
				for _, batch := range batches {
					block := dhtnode.PreCreateNextBlock(batch)
					//将PreCreateNextBlock传给MainNode
					err := dhtnode.DhtInsideTransBlock(block)
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
					timer = time.After(BATCHTIMEOUT)
					logger.Debugf("Just began %s batch timer", BATCHTIMEOUT.String())
				default:
					// Do nothing when:
					// 1. Timer is already running and there are messages pending
					// 2. Timer is not set and there are no messages pending
				}
			} else {
				batch := dhtnode.Cut()
				if batch != nil {
					block := dhtnode.PreCreateNextBlock(batch)
					err := dhtnode.DhtInsideTransBlock(&cb.Block{Header: block.Header, Data: block.Data, Metadata: block.Metadata /*参数v*/})
					if err != nil {
						log.Fatalf("could not transcation Block: %v", err)
					}
				}
				block := dhtnode.PreCreateNextBlock([]*cb.Envelope{msg.ConfigMsg})
				err := dhtnode.DhtInsideTransBlock(&cb.Block{Header: block.Header, Data: block.Data, Metadata: block.Metadata /*参数v*/})
				if err != nil {
					log.Fatalf("could not transcation Block: %v", err)
				}
				timer = nil
			}
		case <-timer:
			//clear the timer
			timer = nil
			batch := dhtnode.Cut()
			if len(batch) == 0 {
				logger.Warningf("Batch timer expired with no pending requests, this might indicate a bug")
				continue
			}
			logger.Debugf("Batch timer expired, creating block")
			block := dhtnode.PreCreateNextBlock(batch)
			err := dhtnode.DhtInsideTransBlock(&cb.Block{Header: block.Header, Data: block.Data, Metadata: block.Metadata /*参数v*/})
			if err != nil {
				log.Fatalf("could not transcation Block: %v", err)
			}

		case <-dhtnode.GetShutdownCh():
			logger.Debugf("Exiting")
			return
		}
	}
}
