package service

import (
	"fmt"
	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
	"lark/pkg/proto/pb_cm"
	"sync/atomic"
)

func (s *cloudMessageService) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.consumerGroup.Ready)
	return nil
}
func (s *cloudMessageService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *cloudMessageService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		msg *sarama.ConsumerMessage
		err error
	)
	for {
		select {
		case msg = <-claim.Messages():
			if msg == nil {
				continue
			}
			if err = s.msgHandle[msg.Topic](msg.Value, string(msg.Key)); err != nil {
				continue
			}
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
	return nil
}

func (s *cloudMessageService) MessageHandler(msg []byte, key string) (err error) {
	var (
		req = new(pb_cm.CloudMessageReq)
	)
	proto.Unmarshal(msg, req)
	atomic.AddInt64(&s.msgCount, 1)

	// TODO:离线推送业务
	fmt.Println("离线推送:", s.msgCount, len(req.Member))
	return
}
