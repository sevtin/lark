package service

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"lark/pkg/proto/pb_cm"
	"sync/atomic"
)

func (s *cloudMessageService) MessageHandler(msg []byte, chatId string) (err error) {
	var (
		req = new(pb_cm.CloudMessageReq)
	)
	proto.Unmarshal(msg, req)
	atomic.AddInt64(&s.msgCount, 1)

	// TODO:离线推送业务
	fmt.Println("离线推送:", s.msgCount, len(req.Member))
	return
}
