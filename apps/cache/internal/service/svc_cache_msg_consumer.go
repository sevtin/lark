package service

import (
	"github.com/Shopify/sarama"
	"lark/domain/do"
	"lark/pkg/constant"
	"lark/pkg/utils"
)

func (s *cacheService) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.consumerGroup.Ready)
	return nil
}
func (s *cacheService) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *cacheService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
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

func (s *cacheService) MessageHandler(msg []byte, key string) (err error) {
	if len(msg) == 0 {
		return
	}
	switch key {
	case constant.CONST_MSG_KEY_CACHE_ON_OFF_LINE:
		var (
			obj = new(do.KeysFieldValues)
		)
		utils.ByteToObj(msg, obj)
		err = s.chatMemberCache.HSetDistChatMembers(obj.Keys, obj.Field, obj.Values)
	case constant.CONST_MSG_KEY_CACHE_AGREE_INVITATION:
		var (
			obj    = new(do.KeyMaps)
			chatId int64
		)
		utils.ByteToObj(msg, obj)
		chatId, _ = utils.ToInt64(obj.Key)
		err = s.chatMemberCache.HMSetChatMembers(chatId, obj.Maps)
	case constant.CONST_MSG_KEY_CACHE_REMOVE_CHAT_MEMBER:
		var (
			obj = new(do.KeysValues)
		)
		utils.ByteToObj(msg, obj)
		err = s.chatMemberCache.HDelChatMembers(obj.Keys, obj.Values)
	case constant.CONST_MSG_KEY_CACHE_CREATE_GROUP_CHAT:
		var (
			obj    = new(do.KeyFieldValue)
			chatId int64
			uid    int64
			value  string
		)
		utils.ByteToObj(msg, obj)
		chatId, _ = utils.ToInt64(obj.Key)
		uid, _ = utils.ToInt64(obj.Field)
		value = utils.ToString(obj.Value)
		err = s.chatMemberCache.HSetNXChatMember(chatId, uid, value)
	}
	return
}
