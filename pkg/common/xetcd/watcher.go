package xetcd

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"lark/pkg/common/xlog"
	"sync"
	"time"
)

type Watcher struct {
	rwLock   sync.RWMutex
	client   *clientv3.Client
	kv       clientv3.KV
	watcher  clientv3.Watcher
	catalog  string
	schema   string
	kvs      map[string]*KeyValue
	address  []string
	callback func(keyValue *KeyValue, eventType int)
	revision int64
}

const (
	EVENT_TYPE_PUT    = 0
	EVENT_TYPE_DELETE = 1
)

type KeyValue struct {
	Key     string
	Value   string
	Version int64
}

func NewWatcher(catalog string, schema string, address []string, callback func(keyValue *KeyValue, eventType int)) (w *Watcher, err error) {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		watcher clientv3.Watcher
	)

	config = clientv3.Config{
		Endpoints:   address,                                // 集群地址
		DialTimeout: time.Duration(5000) * time.Millisecond, // 连接超时
	}
	// 1、建立连接
	if client, err = clientv3.New(config); err != nil {
		xlog.Error(err.Error())
		return
	}
	// 2、得到KV和观察者
	kv = clientv3.NewKV(client)
	watcher = clientv3.NewWatcher(client)

	w = &Watcher{
		client:   client,
		kv:       kv,
		watcher:  watcher,
		catalog:  catalog,
		kvs:      make(map[string]*KeyValue),
		schema:   schema,
		address:  address,
		callback: callback,
	}
	return
}

// 获取当前kvs
func (w *Watcher) Run() (err error) {
	var (
		resp   *clientv3.GetResponse
		kvpair *mvccpb.KeyValue
	)
	// 1、get目录下的所有键值对，并且获知当前集群的revision
	if resp, err = w.kv.Get(context.TODO(), w.catalog, clientv3.WithPrefix()); err != nil {
		xlog.Error(err.Error())
		return
	}
	for _, kvpair = range resp.Kvs {
		kv := &KeyValue{
			Key:     string(kvpair.Key),
			Value:   string(kvpair.Value),
			Version: kvpair.Version,
		}
		w.kvs[kv.Key] = kv
		w.callback(kv, EVENT_TYPE_PUT)
	}
	w.revision = resp.Header.Revision
	return
}

// 监听变化
func (w *Watcher) Watch() {
	var (
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
		key                string
		kv1                *KeyValue
		kv2                *KeyValue
		ok                 bool
		put                bool
	)
	// 2、从该revision向后监听变化事件
	go func() {
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = w.revision + 1
		// 监听目录的后续变化
		watchChan = w.watcher.Watch(context.TODO(), w.catalog, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					kv1 = &KeyValue{
						Key:     string(watchEvent.Kv.Key),
						Value:   string(watchEvent.Kv.Value),
						Version: watchEvent.Kv.Version,
					}
					put = true
					w.rwLock.Lock()
					if kv2, ok = w.kvs[kv1.Key]; ok == true {
						if kv1.Version > kv2.Version {
							put = true
						}
					}
					if put == true {
						w.kvs[kv1.Key] = kv1
					}
					w.rwLock.Unlock()
					if put == true {
						w.callback(kv1, EVENT_TYPE_PUT)
					}
				case mvccpb.DELETE: // 任务被删除了
					key = string(watchEvent.Kv.Key)
					w.rwLock.Lock()
					if kv1, ok = w.kvs[key]; ok == true {
						delete(w.kvs, key)
					}
					w.rwLock.Unlock()
					if kv1 != nil {
						w.callback(kv1, EVENT_TYPE_DELETE)
					}
				}
			}
		}
	}()
}

func (w *Watcher) EachKvs(f func(k string, v *KeyValue) bool) {
	//w.rwLock.RLock()
	//defer w.rwLock.RUnlock()
	for k, v := range w.kvs {
		if !f(k, v) {
			return
		}
	}
	return
}
