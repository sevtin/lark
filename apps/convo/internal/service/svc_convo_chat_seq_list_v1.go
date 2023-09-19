package service

import (
	"context"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_convo"
	"lark/pkg/proto/pb_enum"
)

func (s *convoService) ConvoChatSeqList(ctx context.Context, req *pb_convo.ConvoChatSeqListReq) (resp *pb_convo.ConvoChatSeqListResp, _ error) {
	resp = &pb_convo.ConvoChatSeqListResp{List: make([]*pb_convo.ConvoChatSeq, 0)}
	var (
		q = entity.NewNormalQuery()
	)
	q.SetFilter("m.uid=?", req.Uid)
	q.SetFilter("m.chat_id!=?", req.LastCid)
	q.SetFilter("m.status<=?", pb_enum.CHAT_STATUS_BANNED)
	q.SetFilter("m.deleted_ts=?", 0)
	q.SetFilter("c.deleted_ts=?", 0)
	q.SetFilter("c.srv_ts<=?", req.LastTs)
	q.SetLimit(req.Limit)
	resp.List, _ = s.chatMemberRepo.ConvoChatSeqList(q)
	return
}
