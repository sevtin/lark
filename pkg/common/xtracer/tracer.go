package xtracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"lark/pkg/conf"
)

/*
通过 gRPC-Jaeger 拦截器上报
https://cloud.tencent.com/document/product/1463/57865

http://127.0.0.1:16686/search
*/

func NewTracer(serviceName string, c *conf.Jaeger) (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := &jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  c.SamplerType,
			Param: c.Param,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           c.LogSpans,
			LocalAgentHostPort: c.HostPort,
		},
	}
	tracer, closer, err = cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	return
}
