package xgrpc

import (
	"github.com/opentracing/opentracing-go"
	"io"
	"lark/pkg/common/xtracer"
	"lark/pkg/conf"
)

type ClientDialOption struct {
	DialOption *conf.GrpcDialOption
	closer     io.Closer
}

func NewClientDialOption(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) *ClientDialOption {
	var (
		tracer opentracing.Tracer
		closer io.Closer
		opt    *conf.GrpcDialOption
	)
	if jaeger.Enabled == true {
		tracer, closer, _ = xtracer.NewTracer(clientName, jaeger)
	}
	opt = &conf.GrpcDialOption{
		ServiceName: server.Name,
		Etcd:        etcd,
		Tracing:     &conf.Tracing{Tracer: tracer, Enabled: jaeger.Enabled},
		Cert:        server.Cert,
	}
	return &ClientDialOption{
		DialOption: opt,
		closer:     closer,
	}
}
