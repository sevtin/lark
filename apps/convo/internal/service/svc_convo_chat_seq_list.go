package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_convo"
	"lark/pkg/utils"
	"strings"
)

func (s *convoService) ConvoChatSeqList(ctx context.Context, req *pb_convo.ConvoChatSeqListReq) (resp *pb_convo.ConvoChatSeqListResp, _ error) {
	resp = &pb_convo.ConvoChatSeqListResp{List: make([]*pb_convo.ConvoChatSeq, 0)}
	var (
		buf         []byte
		val         string
		chatIds     []string
		seqIdTsList []string
		index       int
		value       string
		arr         []string
		timestamp   int64
		seq         *pb_convo.ConvoChatSeq
		err         error
	)
	buf, err = utils.DecodeString(req.ChatIds)
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
	if len(chatIds) == 0 {
		return
	}
	if len(chatIds) > MAXIMUM_NUMBER_OF_CHATS {
		resp.Set(ERROR_CODE_CONVO_PARAM_ERR, ERROR_CONVO_PARAM_ERR)
		return
	}
	seqIdTsList, err = s.convoCache.MGetSeqIdTsList(chatIds)
	if err != nil {
		resp.Set(ERROR_CODE_CONVO_REDIS_GET_FAILED, ERROR_CONVO_REDIS_GET_FAILED)
		xlog.Warn(ERROR_CODE_CONVO_REDIS_GET_FAILED, ERROR_CONVO_REDIS_GET_FAILED, err.Error())
		return
	}
	if len(chatIds) != len(seqIdTsList) {
		return
	}
	resp.List = make([]*pb_convo.ConvoChatSeq, 0)
	for index, value = range seqIdTsList {
		if value == "" {
			continue
		}
		arr = strings.Split(value, ",")
		if len(arr) == 2 {
			timestamp, _ = utils.ToInt64(arr[1])
			if timestamp > req.Timestamp {
				seq = new(pb_convo.ConvoChatSeq)
				seq.ChatId, _ = utils.ToInt64(chatIds[index])
				seq.SeqId, _ = utils.ToInt64(arr[0])
				seq.SrvTs = timestamp
				resp.List = append(resp.List, seq)
			}
		}
	}
	return
}
