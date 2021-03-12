package dhtnode

import (
	"bytes"
	"crypto/sha256"
	"time"

	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/zebra-uestc/chord/config"
)

/*
PrevBlock生成类
*/

var logger = flogging.MustGetLogger("orderer.consensus.dht")

type OrdererConfigFetcher interface {
	OrdererConfig() (channelconfig.Orderer, bool)
}

// CreateNextBlock creates a new block with the next block number, and the given contents.
func (dhtn *dhtNode) PreCreateNextBlock(messages []*cb.Envelope) *cb.Block {
	// previousBlockHash := protoutil.BlockHeaderHash(bw.lastBlock.Header)

	data := &cb.BlockData{
		Data: make([][]byte, len(messages)),
	}

	var err error
	for i, msg := range messages {
		data.Data[i], err = proto.Marshal(msg)
		if err != nil {
			logger.Panicf("Could not marshal envelope: %s", err)
		}
	}

	// block := protoutil.NewBlock(bw.lastBlock.Header.Number+1, previousBlockHash)
	block := NewBlock(0, []byte{})
	block.Header.DataHash = protoutil.BlockDataHash(data)
	block.Data = data

	return block
}

func BlockDataHash(b *cb.BlockData) []byte {
	sum := sha256.Sum256(bytes.Join(b.Data, nil))
	return sum[:]
}

func NewBlock(seqNum uint64, previousHash []byte) *cb.Block {
	block := &cb.Block{}
	block.Header = &cb.BlockHeader{}
	block.Header.Number = seqNum
	block.Header.PreviousHash = previousHash
	block.Header.DataHash = []byte{}
	block.Data = &cb.BlockData{}

	var metadataContents [][]byte
	for i := 0; i < len(cb.BlockMetadataIndex_name); i++ {
		metadataContents = append(metadataContents, []byte{})
	}
	block.Metadata = &cb.BlockMetadata{Metadata: metadataContents}

	return block
}

// Ordered should be invoked sequentially as messages are ordered
//
// messageBatches length: 0, pending: false
//   - impossible, as we have just received a message
// messageBatches length: 0, pending: true
//   - no batch is cut and there are messages pending
// messageBatches length: 1, pending: false
//   - the message count reaches BatchSize.MaxMessageCount
// messageBatches length: 1, pending: true
//   - the current message will cause the pending batch size in bytes to exceed BatchSize.PreferredMaxBytes.
// messageBatches length: 2, pending: false
//   - the current message size in bytes exceeds BatchSize.PreferredMaxBytes, therefore isolated in its own batch.
// messageBatches length: 2, pending: true
//   - impossible
//
// Note that messageBatches can not be greater than 2.
func (dhtn *dhtNode) Ordered(msg *cb.Envelope) (messageBatches [][]*cb.Envelope, pending bool) {
	if len(dhtn.pendingBatch) == 0 {
		// We are beginning a new batch, mark the time
		dhtn.PendingBatchStartTime = time.Now()
	}

	messageSizeBytes := messageSizeBytes(msg)
	if messageSizeBytes > config.PreferredMaxBytes {
		logger.Debugf("The current message, with %v bytes, is larger than the preferred batch size of %v bytes and will be isolated.", messageSizeBytes, config.PreferredMaxBytes)

		// cut pending batch, if it has any messages
		if len(dhtn.pendingBatch) > 0 {
			messageBatch := dhtn.Cut()
			messageBatches = append(messageBatches, messageBatch)
		}

		// create new batch with single message
		messageBatches = append(messageBatches, []*cb.Envelope{msg})

		// Record that this batch took no time to fill
		// dhtn.Metrics.BlockFillDuration.With("channel", dhtn.ChannelID).Observe(0)

		return
	}

	messageWillOverflowBatchSizeBytes := dhtn.pendingBatchSizeBytes+messageSizeBytes > config.PreferredMaxBytes

	if messageWillOverflowBatchSizeBytes {
		logger.Debugf("The current message, with %v bytes, will overflow the pending batch of %v bytes.", messageSizeBytes, dhtn.pendingBatchSizeBytes)
		logger.Debugf("Pending batch would overflow if current message is added, cutting batch now.")
		messageBatch := dhtn.Cut()
		dhtn.PendingBatchStartTime = time.Now()
		messageBatches = append(messageBatches, messageBatch)
	}

	logger.Debugf("Enqueuing message into batch")
	dhtn.pendingBatch = append(dhtn.pendingBatch, msg)
	dhtn.pendingBatchSizeBytes += messageSizeBytes
	pending = true

	if uint32(len(dhtn.pendingBatch)) >= config.MaxMessageCount {
		logger.Debugf("Batch size met, cutting batch")
		messageBatch := dhtn.Cut()
		messageBatches = append(messageBatches, messageBatch)
		pending = false
	}

	return
}

// Cut returns the current batch and starts a new one
func (dhtn *dhtNode) Cut() []*cb.Envelope {
	// if dhtn.pendingBatch != nil {
		// dhtn.Metrics.BlockFillDuration.With("channel", dhtn.ChannelID).Observe(time.Since(dhtn.PendingBatchStartTime).Seconds())
	// }
	dhtn.PendingBatchStartTime = time.Time{}
	batch := dhtn.pendingBatch
	dhtn.pendingBatch = nil
	dhtn.pendingBatchSizeBytes = 0
	return batch
}

func messageSizeBytes(message *cb.Envelope) uint32 {
	return uint32(len(message.Payload) + len(message.Signature))
}

