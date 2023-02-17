package main

import (
	"flag"
	"lark/examples/chat/client"
	"lark/examples/config"
	"lark/pkg/common/xredis"
	"sync"
)

/*
------ 单机群聊压测 ------
系统: MacOS 10.15.7
CPU: 3.2 GHz 六核Intel Core i7
内存: 16G
群成员数: 10000人
在线人数: 1000人
每秒/次发送消息数量: 500条
发送消息次数: 400次
转发/接收消息数量: 200000000条
响应消息数量: 200000条
丢失消息数量: 0条
总耗时: 1813317ms
平均每500条消息发送/转发在线人员/在线人员接收总耗时: 4533ms
平均每1条消息发送/转发在线人员/在线人员接收总耗时: 42ms
*/
var (
	// 在线成员数量
	on = flag.Int64("on", 1000, "Online users")
	// 每次发送消息数量
	sn = flag.Int64("sn", 500, "Number of messages send")
	// 群成员数量
	gn = flag.Int64("gn", 10000, "Number of group members")
	// 发送消息次数
	tn = flag.Int64("tn", 400, "Number of tests")
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
