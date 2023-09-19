package pdo

import "lark/pkg/utils"

var (
	field_tag_convo_chat_seq string
)

type ConvoChatSeq struct {
	ReadSeq int   `json:"read_seq" field:"o.read_seq"`
	ChatId  int64 `json:"chat_id" field:"c.chat_id"`
	SeqId   int   `json:"seq_id" field:"c.seq_id"`
	SrvTs   int64 `json:"srv_ts" field:"c.srv_ts"`
}

func (p *ConvoChatSeq) GetFields() string {
	if field_tag_convo_chat_seq == "" {
		field_tag_convo_chat_seq = utils.GetFields(*p)
	}
	return field_tag_convo_chat_seq
}
