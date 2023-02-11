package xgrpc

import (
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
	"lark/pkg/common/xetcd"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/utils"
	"net"
	"strconv"
)

type GrpcServer struct {
	grpc *conf.Grpc
	etcd *conf.Etcd
}

func NewGrpcServer(grpc *conf.Grpc, etcd *conf.Etcd) *GrpcServer {
	return &GrpcServer{grpc, etcd}
}

func (s *GrpcServer) RunServer(server *grpc.Server) {
	var (
		address  string
		listener net.Listener
		err      error
	)
	defer func() {
		server.GracefulStop()
	}()

	address = "0.0.0.0:" + strconv.Itoa(s.grpc.Port)
	listener, err = net.Listen("tcp", address)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	if s.grpc.ConnectionLimit > 0 {
		listener = netutil.LimitListener(listener, s.grpc.ConnectionLimit)
	}
	err = xetcd.RegisterEtcd(s.etcd.Schema, s.etcd.Endpoints, utils.GetServerIP(), s.grpc.Port, s.grpc.Name, 10)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	err = server.Serve(listener)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
}
