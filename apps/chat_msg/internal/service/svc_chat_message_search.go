package service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"lark/pkg/common/xes"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_chat_member"
	"lark/pkg/proto/pb_chat_msg"
	"lark/pkg/utils"
)

func (s *chatMessageService) SearchMessage(ctx context.Context, req *pb_chat_msg.SearchMessageReq) (resp *pb_chat_msg.SearchMessageResp, _ error) {
	resp = &pb_chat_msg.SearchMessageResp{List: make([]*pb_chat_msg.MessageSummary, 0)}
	var (
		buf    bytes.Buffer
		err    error
		res    *esapi.Response
		client *elasticsearch.Client
		r      map[string]interface{}
		i      int
		hit    interface{}
		total  interface{}
	)

	if s.memberVerification(req.ChatId, req.Uid) == false {
		resp.Set(ERROR_CODE_CHAT_MSG_NON_CHAT_MEMBER, ERROR_CHAT_MSG_NON_CHAT_MEMBER)
		return
	}

	query := map[string]interface{}{
		"size": req.Size,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"term": map[string]interface{}{
						"body": req.Query,
					},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"srv_msg_id": map[string]interface{}{
							"gt": req.LastMsgId,
						},
					},
				},
			},
		},
		"sort": []map[string]interface{}{
			{"srv_msg_id": "desc"},
		},
		"highlight": map[string]interface{}{
			"pre_tags":  "<b class='key' style='color:red'>",
			"post_tags": "</b>",
			"fields": map[string]interface{}{
				"body": map[string]interface{}{},
			},
		},
	}

	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		resp.Set(ERROR_CODE_CHAT_MSG_ENCODING_FAILED, ERROR_CHAT_MSG_ENCODING_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	client = xes.GetClient()
	res, err = client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("messages"),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MSG_SEARCH_FAILED, ERROR_CHAT_MSG_SEARCH_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	defer res.Body.Close()
	if res.IsError() {
		resp.Set(ERROR_CODE_CHAT_MSG_SEARCH_FAILED, ERROR_CHAT_MSG_SEARCH_FAILED)
		xlog.Warn(resp.Code, resp.Msg, res.StatusCode)
		return
	}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		resp.Set(ERROR_CODE_CHAT_MSG_DECODE_FAILED, ERROR_CHAT_MSG_DECODE_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	resp.List = make([]*pb_chat_msg.MessageSummary, len(r["hits"].(map[string]interface{})["hits"].([]interface{})))
	total = r["hits"].(map[string]interface{})["total"]
	resp.Total, _ = utils.ToInt64(total.(map[string]interface{})["value"])

	for i, hit = range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var msg = &pb_chat_msg.MessageSummary{}
		utils.Copy(hit.(map[string]interface{})["_source"], msg)
		highlight := hit.(map[string]interface{})["highlight"]
		msg.Rt = highlight.(map[string]interface{})["body"].([]interface{})[0].(string)
		resp.List[i] = msg
	}
	return
}

func (s *chatMessageService) memberVerification(chatId int64, uid int64) (ok bool) {
	var (
		memberInfo *pb_chat_member.ChatMemberInfo
		reqArgs    = &pb_chat_member.GetChatMemberInfoReq{
			ChatId: chatId,
			Uid:    uid,
		}
		reply *pb_chat_member.GetChatMemberInfoResp
		err   error
	)
	memberInfo, err = s.chatMemberCache.GetChatMemberInfo(chatId, uid)
	if memberInfo.Uid > 0 {
		ok = true
		return
	}
	if err != nil {
		xlog.Warn(ERROR_CODE_CHAT_MSG_REDIS_GET_FAILED, ERROR_CHAT_MSG_REDIS_GET_FAILED, err.Error())
		err = nil
	}
	reply = s.chatMemberClient.GetChatMemberInfo(reqArgs)
	if reply == nil {
		xlog.Warn(ERROR_CODE_CHAT_MSG_GRPC_SERVICE_FAILURE, ERROR_CHAT_MSG_GRPC_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		xlog.Warn(reply.Code, reply.Msg)
		return
	}
	if reply.Info.Uid > 0 {
		ok = true
		return
	}
	return
}
