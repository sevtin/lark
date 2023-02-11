package gateway

import (
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/proto/pb_obj"
)

func (s *gatewayServer) sendOperationMessage(req *pb_gw.SendTopicMessageReq, resp *pb_gw.SendTopicMessageResp) {
	var (
		buf []byte
		err error
	)
	buf, err = Encode(req.Topic, req.SubTopic, pb_enum.MESSAGE_TYPE_NEW, req.Body, resp)
	if err != nil {
		return
	}
	var (
		member   *pb_obj.Int64Array
		uid      int64
		platform int32
	)
	// 0:ServerId, 1:Platform, 2:Uid, 3:Status
	for _, member = range req.Members {
		uid = member.GetUid()
		platform = int32(member.GetPlatform())
		s.wsServer.SendMessage(uid, platform, buf)
	}
}
