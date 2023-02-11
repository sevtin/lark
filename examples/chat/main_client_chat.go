package main

import (
	"flag"
	"lark/examples/chat/client"
	"lark/examples/config"
	"lark/pkg/common/xredis"
	"sync"
)

var (
	// 在线成员数量
	on = flag.Int64("on", 3000, "Online users")
	// 每次发送消息数量
	sn = flag.Int64("sn", 1000, "Number of messages send")
	// 群成员数量
	gn = flag.Int64("gn", 10000, "Number of group members")
	// 发送消息次数
	tn = flag.Int64("tn", 1000000000, "Number of tests")
	// chat id
	cid = flag.Int64("cid", 3333336666669999990, "Test Chat ID")
	// 是否集群
	cl = flag.Bool("cl", false, "Is it a cluster")
	// 消息网关数量(需要先部署)
	sc = flag.Int("sc", 3, "Number of Servers")
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(1)

	manager := client.NewManager(*on, *sn, *gn, *tn, *cid, *cl, *sc)
	manager.Run()

	wg.Wait()
}
