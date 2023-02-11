package service

import (
	"context"
	"encoding/base64"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_convo"
	"lark/pkg/utils"
	"strings"
)

func (s *convoService) ConvoChatSeqList(ctx context.Context, req *pb_convo.ConvoChatSeqListReq) (resp *pb_convo.ConvoChatSeqListResp, _ error) {
	resp = &pb_convo.ConvoChatSeqListResp{List: make([]*pb_convo.ConvoChatSeq, 0)}
	var (
		buf       []byte
		val       string
		chatIds   []string
		seqIdList []interface{}
		chatId    string
		index     int
		seq       *pb_convo.ConvoChatSeq
		err       error
	)
	buf, err = base64.StdEncoding.DecodeString(req.ChatIds)
	if err != nil {
		resp.Set(ERROR_CODE_CONVO_DECODE_FAILED, ERROR_CONVO_DECODE_FAILED)
		xlog.Warn(ERROR_CODE_CONVO_DECODE_FAILED, ERROR_CONVO_DECODE_FAILED, err.Error())
		return
	}
	buf, err = utils.UnGzip(buf)
	if err != nil {
		resp.Set(ERROR_CODE_CONVO_UN_GZIP_FAILED, ERROR_CONVO_UN_GZIP_FAILED)
		xlog.Warn(ERROR_CODE_CONVO_UN_GZIP_FAILED, ERROR_CONVO_UN_GZIP_FAILED, err.Error())
		return
	}
	val = string(buf)
	chatIds = strings.Split(val, ",")
	seqIdList, err = s.convoCache.MGetSeqIdList(s.cfg.Redis.Prefix, chatIds)
	if err != nil {
		resp.Set(ERROR_CODE_CONVO_REDIS_GET_FAILED, ERROR_CONVO_REDIS_GET_FAILED)
		xlog.Warn(ERROR_CODE_CONVO_REDIS_GET_FAILED, ERROR_CONVO_REDIS_GET_FAILED, err.Error())
		return
	}
	if len(seqIdList) == 0 {
		return
	}
	if len(chatIds) != len(seqIdList) {
		return
	}
	resp.List = make([]*pb_convo.ConvoChatSeq, len(chatIds))
	for index, chatId = range chatIds {
		seq = new(pb_convo.ConvoChatSeq)
		seq.ChatId, _ = utils.ToInt64(chatId)
		if seqIdList[index] != nil {
			seq.SeqId, _ = utils.ToInt64(seqIdList[index])
		}
		resp.List[index] = seq
	}
	return
}
