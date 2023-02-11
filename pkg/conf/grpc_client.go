package conf

import "github.com/opentracing/opentracing-go"

type GrpcDialOption struct {
	ServiceName string
	Etcd        *Etcd
	Tracing     *Tracing
	Cert        *Cert
}

type Tracing struct {
	Tracer  opentracing.Tracer
	Enabled bool `yaml:"enabled"`
}
