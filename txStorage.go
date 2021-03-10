package chord

import (
	"errors"
	"hash"

	bm "github.com/zebra-uestc/chord/models/bridge"
	cm "github.com/zebra-uestc/chord/models/chord"
	"google.golang.org/protobuf/proto"
)

var emptyMethodError = errors.New("Not Implemented Method")

type TxStorage interface {
	Storage
	GetMsgChan() chan *message
}

func NewtxStorage(hashFunc func() hash.Hash) TxStorage {
	return &txStorage{}
}

type txStorage struct {
	setMsgChan chan *message
}

type message struct {
	ConfigSeq uint64
	NormalMsg *bm.Envelope
	ConfigMsg *bm.Envelope
}

func (txs *txStorage) Set(key []byte, value []byte) error {
	var msgByte *bm.Msg
	msg := &message{}
	if err := proto.Unmarshal(value, msgByte); err != nil{
		return err
	}
	if err := proto.Unmarshal(msgByte.NormalMsg, msg.NormalMsg); err != nil {
		return err
	}
	if err := proto.Unmarshal(msgByte.ConfigMsg, msg.ConfigMsg); err != nil {
		return err
	}
	msg = &message{ConfigSeq:msgByte.ConfigSeq, NormalMsg:msg.NormalMsg, ConfigMsg:msg.ConfigMsg}
	txs.setMsgChan <- msg
	return nil
}

func (txs *txStorage) GetMsgChan() chan *message {
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
