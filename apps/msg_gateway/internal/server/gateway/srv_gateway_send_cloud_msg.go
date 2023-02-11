package gateway

import (
	"fmt"
	"lark/pkg/proto/pb_cm"
)

func (s *gatewayServer) SendCloudMessage(req *pb_cm.CloudMessageReq) {
	s.producer.EnQueue(req, fmt.Sprintf("%d:%d", req.Topic, req.SubTopic))
}
