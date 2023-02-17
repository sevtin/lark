package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_convo"
	"lark/pkg/utils"
	"strings"
)

func (s *convoService) ConvoList(ctx context.Context, req *pb_convo.ConvoListReq) (resp *pb_convo.ConvoListResp, _ error) {
	resp = &pb_convo.ConvoListResp{List: make([]*pb_convo.ConvoMessage, 0)}
	var (
		buf    []byte
		cidStr string
		cids   []string
		keys   []string
		index  int
		cid    string
		key    string
		list   []interface{}
		val    interface{}
		msg    *pb_convo.ConvoMessage
		err    error
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
	cidStr = string(buf)
	cids = strings.Split(cidStr, ",")
	if len(cids) == 0 {
		return
	}
	if len(cids) > MAXIMUM_NUMBER_OF_CONVERSATIONS {
		resp.Set(ERROR_CODE_CONVO_PARAM_ERR, ERROR_CONVO_PARAM_ERR)
		return
	}
	keys = make([]string, len(cids))
	for index, cid = range cids {
		key = s.cfg.Redis.Prefix + constant.RK_MSG_CONVO_MSG + cid
		keys[index] = key
	}
	list, err = s.chatMessageCache.MGetMessages(keys...)
	if err != nil {
		resp.Set(ERROR_CODE_CONVO_REDIS_GET_FAILED, ERROR_CONVO_REDIS_GET_FAILED)
		xlog.Warn(ERROR_CODE_CONVO_REDIS_GET_FAILED, ERROR_CONVO_REDIS_GET_FAILED, err.Error())
		return
	}
	for index, val = range list {
		if val == nil {
			continue
		}
		msg = new(pb_convo.ConvoMessage)
		utils.Unmarshal(val.(string), msg)
		resp.List = append(resp.List, msg)
	}
	return
}
