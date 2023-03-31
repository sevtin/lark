package service

import (
	"context"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_convo"
	"lark/pkg/utils"
	"sort"
	"strings"
)

func (s *convoService) ConvoList(ctx context.Context, req *pb_convo.ConvoListReq) (resp *pb_convo.ConvoListResp, _ error) {
	resp = &pb_convo.ConvoListResp{List: make([]*pb_convo.ConvoMessage, 0)}
	var (
		buf    []byte
		cidStr string
		cids   []string
		maps   = map[uint16][]string{}
		crc16  uint16
		cid    string
		cidVal int64
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
	for _, cid = range cids {
		cidVal, _ = utils.ToInt64(cid)
		key = xredis.GetPrefix() + constant.RK_MSG_CONVO_MSG + utils.GetHashTagKey(cidVal)
		crc16 = utils.GetCrc16(cidVal)
		maps[crc16] = append(maps[crc16], key)
	}
	list, err = s.chatMessageCache.SlotMGetMessages(maps)
	if err != nil {
		resp.Set(ERROR_CODE_CONVO_REDIS_GET_FAILED, ERROR_CONVO_REDIS_GET_FAILED)
		xlog.Warn(ERROR_CODE_CONVO_REDIS_GET_FAILED, ERROR_CONVO_REDIS_GET_FAILED, err.Error())
		return
	}
	for _, val = range list {
		if val == nil {
			continue
		}
		msg = new(pb_convo.ConvoMessage)
		utils.Unmarshal(val.(string), msg)
		resp.List = append(resp.List, msg)
	}
	sort.Slice(resp.List, func(i, j int) bool {
		return resp.List[i].SrvTs < resp.List[j].SrvTs
	})
	return
}
