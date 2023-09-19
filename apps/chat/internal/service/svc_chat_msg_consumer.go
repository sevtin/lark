package service

import (
	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_msg"
)

// Setup is run at the beginning of a new session, before ConsumeClaim
func (s *chatService) Setup(_ sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(s.consumerGroup.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (s *chatService) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (s *chatService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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

func (s *chatService) MessageHandler(msg []byte, msgKey string) (err error) {
	switch msgKey {
	case constant.CONST_MSG_KEY_READ_RECEIPT:
		s.updateReadReceiptSeq(msg)
		return
	case constant.CONST_MSG_KEY_CHAT_SEQ:
		s.updateChatSeq(msg)
		return
	default:
		return
	}
}

func (s *chatService) updateReadReceiptSeq(msg []byte) (err error) {
	var (
		read = new(pb_msg.ReadReceipt)
		u    = entity.NewMysqlUpdate()
	)
	if err = proto.Unmarshal(msg, read); err != nil {
		err = nil
		return
	}
	u.SetFilter("chat_id=?", read.ChatId)
	u.SetFilter("uid=?", read.Uid)
	u.Set("read_seq", read.SeqId)
	s.chatMemberRepo.UpdateChatMember(u)
	return
}

func (s *chatService) updateChatSeq(msg []byte) (err error) {
	var (
		read = new(pb_msg.ChatSeq)
		u    = entity.NewMysqlUpdate()
	)
	if err = proto.Unmarshal(msg, read); err != nil {
		err = nil
		return
	}
	u.SetFilter("chat_id=?", read.ChatId)
	u.Set("seq_id", read.SeqId)
	u.Set("srv_ts", read.SrvTs)
	s.chatRepo.UpdateChat(u)
	return
}
