package xetcd

import (
	"context"
	"fmt"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"strings"
	"sync"
	"time"
)

const (
	CONST_DURATION_GRPC_TIMEOUT_SECOND = 5 * time.Second
)

type Resolver struct {
	opt                *conf.GrpcDialOption
	cc                 resolver.ClientConn
	schema             string
	serviceName        string
	grpcClientConn     *grpc.ClientConn
	cli                *clientv3.Client
	watchStartRevision int64
}

var (
	resolvers     = make(map[string]*Resolver)
	resolverMutex sync.RWMutex
)

func NewResolver(opt *conf.GrpcDialOption) (r *Resolver, err error) {
	var (
		cli *clientv3.Client
	)

	cli, err = clientv3.New(clientv3.Config{
		Endpoints: opt.Etcd.Endpoints,
		Username:  opt.Etcd.Username,
		Password:  opt.Etcd.Password,
	})
	if err != nil {
		xlog.Error(err.Error())
		return
	}

	r = new(Resolver)
	r.opt = opt
	r.schema = opt.Etcd.Schema
	r.serviceName = opt.ServiceName
	r.cli = cli
	resolverMutex.Lock()
	// concurrent map writes
	// Everything throw does should be recursively nosplit so it can be called even when it's unsafe to grow the stack.
	resolver.Register(r)
	r.grpcClientConn, _ = r.newGrpcClientConn()
	resolverMutex.Unlock()
	return
}

func (r *Resolver) newGrpcClientConn() (conn *grpc.ClientConn, err error) {
	var (
		opts []grpc.DialOption
		ctx  context.Context
	)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if r.opt.Tracing.Enabled == true && r.opt.Tracing.Tracer != nil {
		opts = append(opts, grpc.WithBlock())
		opts = append(opts, grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(r.opt.Tracing.Tracer)))
	}
	if r.opt.Cert.Enabled == true {
		var creds credentials.TransportCredentials
		creds, err = credentials.NewClientTLSFromFile(r.opt.Cert.CertFile, r.opt.Cert.ServerNameOverride)
		if err != nil {
			xlog.Error(err.Error())
		} else {
			opts = append(opts, grpc.WithTransportCredentials(creds))
		}
	}
	ctx, _ = context.WithTimeout(context.Background(), CONST_DURATION_GRPC_TIMEOUT_SECOND)
	conn, err = grpc.DialContext(ctx, GetPrefix(r.opt.Etcd.Schema, r.opt.ServiceName), opts...)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	return
}

func (r *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {
}

func (r *Resolver) Close() {
}

func GetConn(opt *conf.GrpcDialOption) *grpc.ClientConn {
	var (
		r   *Resolver
		key = opt.Etcd.Schema + opt.ServiceName
		ok  bool
		err error
	)
	resolverMutex.RLock()
	r, ok = resolvers[key]
	resolverMutex.RUnlock()
	if ok == true {
		return r.grpcClientConn
	}

	r, err = NewResolver(opt)
	if err != nil {
		return nil
	}
	resolverMutex.Lock()
	resolvers[key] = r
	resolverMutex.Unlock()
	return r.grpcClientConn
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	if r.cli == nil {
		return nil, fmt.Errorf("etcd clientv3 client failed, etcd:%s", target)
	}
	r.cc = cc
	ctx, _ := context.WithTimeout(context.Background(), CONST_DURATION_GRPC_TIMEOUT_SECOND)
	prefix := GetPrefix(r.schema, r.serviceName)
	resp, err := r.cli.Get(ctx, prefix, clientv3.WithPrefix())
	if err == nil {
		var addrList []resolver.Address
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: string(resp.Kvs[i].Value)})
		}
		r.cc.UpdateState(resolver.State{Addresses: addrList})
		r.watchStartRevision = resp.Header.Revision + 1
		go r.watch(prefix, addrList)
	} else {
		return nil, fmt.Errorf("etcd get failed, prefix: %s", prefix)
	}
	return r, nil
}

func (r *Resolver) Scheme() string {
	return r.schema
}

func exists(addrList []resolver.Address, addr string) bool {
	for _, v := range addrList {
		if v.Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func (r *Resolver) watch(prefix string, addrList []resolver.Address) {
	rch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for n := range rch {
		update := false
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, string(ev.Kv.Value)) {
					update = true
					addrList = append(addrList, resolver.Address{Addr: string(ev.Kv.Value)})
				}
			case mvccpb.DELETE:
				i := strings.LastIndexAny(string(ev.Kv.Key), "/")
				if i < 0 {
					return
				}
				t := string(ev.Kv.Key)[i+1:]
				if s, ok := remove(addrList, t); ok {
					update = true
					addrList = s
				}
			}
		}

		if update == true {
			r.cc.UpdateState(resolver.State{Addresses: addrList})
		}
	}
}
