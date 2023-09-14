package payment_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_pay"
)

// 弃用 改用 http 服务器
type PaymentClient interface {
	AlipayReturn(req *pb_pay.AlipayReturnReq) (resp *pb_pay.AlipayReturnResp)
	AlipayNotify(req *pb_pay.AlipayNotifyReq) (resp *pb_pay.AlipayNotifyResp)
	PaypalReturn(req *pb_pay.PaypalReturnReq) (resp *pb_pay.PaypalReturnResp)
	PaypalCancel(req *pb_pay.PaypalCancelReq) (resp *pb_pay.PaypalCancelResp)
	PaypalNotify(req *pb_pay.PaypalNotifyReq) (resp *pb_pay.PaypalNotifyResp)
}

type paymentClient struct {
	opt *xgrpc.ClientDialOption
}

func NewPaymentClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) PaymentClient {
	return &paymentClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *paymentClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *paymentClient) AlipayReturn(req *pb_pay.AlipayReturnReq) (resp *pb_pay.AlipayReturnResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_pay.NewPayClient(conn)
	var err error
	resp, err = client.AlipayReturn(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *paymentClient) AlipayNotify(req *pb_pay.AlipayNotifyReq) (resp *pb_pay.AlipayNotifyResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_pay.NewPayClient(conn)
	var err error
	resp, err = client.AlipayNotify(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *paymentClient) PaypalReturn(req *pb_pay.PaypalReturnReq) (resp *pb_pay.PaypalReturnResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_pay.NewPayClient(conn)
	var err error
	resp, err = client.PaypalReturn(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *paymentClient) PaypalCancel(req *pb_pay.PaypalCancelReq) (resp *pb_pay.PaypalCancelResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_pay.NewPayClient(conn)
	var err error
	resp, err = client.PaypalCancel(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *paymentClient) PaypalNotify(req *pb_pay.PaypalNotifyReq) (resp *pb_pay.PaypalNotifyResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_pay.NewPayClient(conn)
	var err error
	resp, err = client.PaypalNotify(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
