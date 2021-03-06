package main

import (
	"log"
	"os"
	"os/signal"

	// "strconv"
	"time"

	"github.com/zebra-uestc/chord/config"
	"github.com/zebra-uestc/chord/dhtnode"
)

var mainNode dhtnode.MainNode

func main() {
	mainNode, err := dhtnode.NewMainNode()
	//mainnode节点启动的地址为8001
	mainNode.StartDht("0", config.MainNodeAddressLocal)
	//mainNode接收来自其他DHT节点的block的地址8002
	mainNode.StartTransBlockServer(config.MainNodeAddressBlock)
	//mainNode接受来自orderer的Msg的地址8003
	mainNode.StartTransMsgServer(config.MainNodeAddressMsg)

	if err != nil {
		log.Fatalln(err)
	}
	//根据主节点创建其他节点，构成哈希环
	// for index, address := range addresses {
	// 	go func (index int,address string)  {
	// 		var err error
	// 		if index == 0 {
	// 			mainNode, err = dhtnode.NewMainNode("0", address)
	// 		} else {
	// 			err = mainNode.AddNode(strconv.Itoa(index+1), address)
	// 		}
	// 		if err != nil {
	// 			log.Fatalln(err)
	// 		}
	// 	}(index,address)
	// }

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-time.After(1 * time.Second)
	<-c
	mainNode.Stop()
}
