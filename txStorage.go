package chord

import (
	"errors"
	"hash"

	bm "github.com/zebra-uestc/chord/models/bridge"
	cb "github.com/hyperledger/fabric-protos-go/common"
	cm "github.com/zebra-uestc/chord/models/chord"
	"github.com/hyperledger/fabric/protoutil"
	"google.golang.org/protobuf/proto"
)

var emptyMethodError = errors.New("Not Implemented Method")

type TxStorage interface {
	Storage
	GetMsgChan() chan *Message
}

func NewtxStorage(hashFunc func() hash.Hash) TxStorage {
	
	return &txStorage{
		setMsgChan: make(chan *Message, 10),
	}
}

type txStorage struct {
	setMsgChan chan *Message
}

type Message struct {
	ConfigSeq uint64
	NormalMsg *cb.Envelope
	ConfigMsg *cb.Envelope
}

func (txs *txStorage) Set(key []byte, value []byte) error {
	msgByte := &bm.MsgBytes{}
	if err := proto.Unmarshal(value, msgByte); err != nil{
		return err
	}
	msg := &Message{
		ConfigSeq: msgByte.ConfigSeq, 
		NormalMsg: protoutil.UnmarshalEnvelopeOrPanic(msgByte.NormalMsg), 
		ConfigMsg: protoutil.UnmarshalEnvelopeOrPanic(msgByte.ConfigMsg),
	}
	txs.setMsgChan <- msg
	return nil
}

func (txs *txStorage) GetMsgChan() chan *Message {
	return txs.setMsgChan
}

func (*txStorage) Get([]byte) ([]byte, error) {
	return nil, emptyMethodError
}
func (*txStorage) Delete([]byte) error {
	return emptyMethodError
}
func (*txStorage) Between([]byte, []byte) ([]*cm.KV, error) {
	return nil, emptyMethodError
}
func (*txStorage) MDelete(...[]byte) error {
	return emptyMethodError
}
