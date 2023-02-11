package convo

import (
	"context"
	"lark/pkg/proto/pb_convo"
)

func (s *convoServer) ConvoList(ctx context.Context, req *pb_convo.ConvoListReq) (resp *pb_convo.ConvoListResp, err error) {
	return s.convoService.ConvoList(ctx, req)
}

func (s *convoServer) ConvoChatSeqList(ctx context.Context, req *pb_convo.ConvoChatSeqListReq) (resp *pb_convo.ConvoChatSeqListResp, err error) {
	return s.convoService.ConvoChatSeqList(ctx, req)
}
