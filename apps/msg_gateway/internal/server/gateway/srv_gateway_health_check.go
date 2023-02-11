package gateway

import (
	"context"
	"lark/pkg/proto/pb_gw"
)

func (s *gatewayServer) HealthCheck(ctx context.Context, req *pb_gw.HealthCheckReq) (resp *pb_gw.HealthCheckResp, err error) {
	resp = new(pb_gw.HealthCheckResp)
	return
}
