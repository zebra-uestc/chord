package main

import (
	"os"
	"os/signal"

	// "strconv"
	"time"

	"github.com/zebra-uestc/chord"
	"github.com/zebra-uestc/chord/dhtnode"
	cm "github.com/zebra-uestc/chord/models/chord"
)


var mainNode dhtnode.MainNode

var addresses = []string{
	"127.0.0.1:7001",
	"127.0.0.1:7002",
	"127.0.0.1:7003",
}



//启动其他节点，id,address为各节点本地addr
func StartDht(id string, address string, joinNode *cm.Node) {
	nodeCnf := chord.DefaultConfig()
	nodeCnf.Id = id
	nodeCnf.Addr = address
	nodeCnf.Timeout = 10 * time.Millisecond
	nodeCnf.MaxIdle = 100 * time.Millisecond
	dhtNode, _ := dhtnode.NewDhtNode(nodeCnf, joinNode)
	dhtNode.IsMainNode = false
}


func main() {
	
	//mainnode节点启动的地址为8001
	sister := chord.NewInode("0", "0.0.0.0:8001")
	StartDht("4", addresses[0], sister)
	StartDht("5", addresses[1], sister)

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
