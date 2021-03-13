package main

import (
	"os"
	"time"
	"flag"
	"os/signal"

	"github.com/zebra-uestc/chord"
	"github.com/zebra-uestc/chord/dhtnode"
	"github.com/zebra-uestc/chord/config"
	cm "github.com/zebra-uestc/chord/models/chord"
)

// 启动其他节点，id,address为各节点本地addr
func startDht(id string, address string, joinNode *cm.Node) {
	nodeCnf := chord.DefaultConfig()
	nodeCnf.Id = id
	nodeCnf.Addr = address
	nodeCnf.Timeout = 10 * time.Millisecond
	nodeCnf.MaxIdle = 100 * time.Millisecond
	dhtNode, _ := dhtnode.NewDhtNode(nodeCnf, joinNode)
	dhtNode.IsMainNode = false
}

func main() {
	var (
		addr string
		id string
	)
	flag.StringVar(&addr, "addr", "", "The node addr.")
	flag.StringVar(&id, "id", "", "The node id.")
	flag.Parse()
	if(addr == ""){
		println("please input the node addr")
		return
	}
	if(id == ""){
		println("please input the node id")
		return
	}
	// 加入到mainnode的环上
	startDht(id, addr, chord.NewInode("0", config.MainNodeAddressLocal))

	// waiting ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
