package lbs_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_lbs"
)

type LbsClient interface {
	ReportLngLat(req *pb_lbs.ReportLngLatReq) (resp *pb_lbs.ReportLngLatResp)
	PeopleNearby(req *pb_lbs.PeopleNearbyReq) (resp *pb_lbs.PeopleNearbyResp)
}

type lbsClient struct {
	opt *xgrpc.ClientDialOption
}

func NewLbsClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) LbsClient {
	return &lbsClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *lbsClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *lbsClient) ReportLngLat(req *pb_lbs.ReportLngLatReq) (resp *pb_lbs.ReportLngLatResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_lbs.NewLbsClient(conn)
	var err error
	resp, err = client.ReportLngLat(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *lbsClient) PeopleNearby(req *pb_lbs.PeopleNearbyReq) (resp *pb_lbs.PeopleNearbyResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_lbs.NewLbsClient(conn)
	var err error
	resp, err = client.PeopleNearby(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
