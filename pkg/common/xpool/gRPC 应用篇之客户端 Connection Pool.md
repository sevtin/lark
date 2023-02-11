https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/

## 0x00 连接池[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#0x00----%E8%BF%9E%E6%8E%A5%E6%B1%A0)

之前分析过 `go-redis` 的连接池实现：[Go-Redis 连接池（Pool）源码分析](https://pandaychen.github.io/2020/02/22/A-REDIS-POOL-ANALYSIS/)，在项目应用中，连接池的最大好处是减少 TCP 建立握手 / 挥手的时间，实现 TCP 连接复用，从而降低通信时延和提高性能。

通常一些高性能中间件都提供了内置的 TCP 连接池，如刚才说的 `go-redis`,[`go-sql-driver`](https://github.com/go-sql-driver/mysql#connection-pool-and-timeouts) 等等，关于连接池，一个良好的设计是对用户屏蔽底层的实现，如存储 / keepalive / 关闭 / 自动重连等，只对用户提供简单的获取接口。

## 0x01 gRPC 连接池的实现[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#0x01--grpc-%E8%BF%9E%E6%8E%A5%E6%B1%A0%E7%9A%84%E5%AE%9E%E7%8E%B0)

#### 实现原则[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#%E5%AE%9E%E7%8E%B0%E5%8E%9F%E5%88%99)

1、连接池的基本属性  
首先要选择合适的数据结构来存放连接池（如slice、map、linklist等），通常连接池属性包含**最大空闲连接数**、**最大活跃连接数**以及**最小活跃连接数**等，定义分别如下：

- 最大空闲连接数：连接池一直保持的连接数，无论这些连接被使用与否都会被保持。如果客户端对连接池的使用量不大，便会造成服务端连接资源的浪费
- 最大活跃连接数：连接池最多保持的连接数，如果客户端请求超过次数，便要根据池满的处理机制来处理没有得到连接的请求
- 最小活跃连接数：连接池初始化或是其他时间，连接池内一定要存储的活跃连接数

2、连接池的扩缩容机制如何实现

- 扩容：当请求到来时，如果连接池中没有空闲的连接，同时连接数也没有达到最大活跃连接数，便会按照特定的增长策略创建新的连接服务该请求，同时**用完之后归还到池中**，而非关闭连接
- 缩容：当连接池一段时间没有被使用，同时池中的连接数超过了最大空闲连接数，那么便会关闭一部分连接，使池中的连接数始终维持在最大空闲连接数

3、空闲连接的超时与keepalive

- 超时：如果连接没有被客户端使用的话，便会成为空闲连接，在一段时间后，服务端可能会根据自己的超时策略关闭空闲连接，此时空闲连接已经失效，如果客户端再使用失效的连接，便会通信失败。为了避免这种情况发生，通常连接池中的连接设有最大空闲超时时间(最好略小于服务器的空闲连接超时时间)，在从池中获取连接时，判断是否空闲超时，如果超时则关闭，没有超时则可以继续使用

- keepalive：一般使用某种机制（如心跳包等）保活某个连接，防止服务端/客户端主动Reset；如果服务器发生重启，那么连接池中的连接便会全部失效，即连接池失效了，如何优化此类场景呢？


如何解决上述场景keepalive的失效问题呢？

1. 连接池设置一个Ping函数，专门用来做连接的保活。在从池中获取连接的时候，Ping一下服务器，如果得到响应，则连接依然有效，便可继续使用，如果超时无响应，则关闭该连接，生成新的连接，由于每次都要Ping一下，必然会增加延迟。也可以后台用一个线程或者协程定期的执行Ping函数，进行连接的保活，缺点是感知连接的失效会有一定的延迟，从池中仍然有可能获取到失效的连接。

2. 客户端加入相应的重试机制。比如重试`3`次，前两次从池中获取连接执行，如果报的错是失效的连接等有关连接问题的错误，那么第3次从池中获取的时候带上参数，指定获取新建的连接，同时连接池移除前两次获取的失效的连接。


4、连接池满的处理机制  
当连接池容量超上限时，有`2`种处理机制：

1. 对于连接池新建的连接，并返回给客户端，当客户端用完时，如果池满则关闭连接，否则放入池中
2. 设置一定的超时时间来等待空闲连接。需要客户端加入重试机制，避免因超时之后获取不到空闲连接产生的错误

5、连接池异常的容错机制  
连接池异常时，退化为新建连接的方式，避免影响正常请求，同时，需要相关告警通知开发人员

6、开启异步连接池回收  
参考go-redis的连接池实现，对于空闲连接（超过允许最大空闲时间）超时后，主动关闭连接池中的连接

#### 连接池实现的步骤[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#%E8%BF%9E%E6%8E%A5%E6%B1%A0%E5%AE%9E%E7%8E%B0%E7%9A%84%E6%AD%A5%E9%AA%A4)

1. 服务启动时建立连接池
2. 初始化连接池，建立最大空闲连接数个连接
3. 请求到来时，从池中获取一个连接。如果没有空闲连接且连接数没有达到最大活跃连接数，则新建连接；如果达到最大活跃连接数，允许设置一定的超时时间，等待获取空闲连接
4. 获取到连接后进行通信服务
5. 释放连接，此时是将连接放回连接池，如果池满则关闭连接
6. 释放连接池，关闭所有连接

#### gRPC的特性[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#grpc%E7%9A%84%E7%89%B9%E6%80%A7)

在实现gRPC的连接池之前，需要先了解gRPC的多路复用及超时重连两个特性：

1、多路复用  
gRPC使用HTTP/2作为应用层的传输协议，HTTP/2会复用底层的TCP连接。每一次RPC调用会产生一个新的Stream，每个Stream包含多个Frame，Frame是HTTP/2里面最小的数据传输单位。同时每个Stream有唯一的ID标识，如果是客户端创建的则ID是奇数，服务端创建的ID则是偶数。如果一条连接上的ID使用完了，Client会新建一条连接，Server也会给Client发送一个`GOAWAY Frame`强制让Client新建一条连接。一条gRPC连接允许并发的发送和接收多个Stream，控制的参数便是`MaxConcurrentStreams`

2、超时重连  
在通过调用`Dial`/`DialContext`方法创建连接时，默认只是返回`ClientConn`结构体指针，同时会启动一个goroutine异步的去建立连接。如果想要等连接建立完再返回，可以指定`grpc.WithBlock()`传入`Options`来实现。

超时机制很简单，在调用的时候传入一个timeout的`context`就可以了。重连机制通过启动一个goroutine异步的去建立连接实现的，可以避免服务器因为连接空闲时间过长关闭连接、服务器重启等造成的客户端连接失效问题。也就是说通过**gRPC的重连机制可以完美的解决连接池设计原则中的空闲连接的超时与keepalive问题**。

3、gRPC默认参数优化（基于大块数据传输场景）

```
MaxSendMsgSizeGRPC	//最大允许发送的字节数，默认4MiB，如果超过了GRPC会报错。Client和Server我们都调到4GiB

MaxRecvMsgSizeGRPC	//最大允许接收的字节数，默认4MiB，如果超过了GRPC会报错。Client和Server我们都调到4GiB

InitialWindowSize	//基于Stream的滑动窗口，类似于TCP的滑动窗口，用来做流控，默认64KiB，吞吐量上不去，Client和Server我们调到1GiB

InitialConnWindowSize	//基于Connection的滑动窗口，默认16 * 64KiB，吞吐量上不去，Client和Server我们也都调到1GiB

KeepAliveTime	//每隔KeepAliveTime时间，发送PING帧测量最小往返时间，确定空闲连接是否仍然有效，我们设置为10s

KeepAliveTimeout	//超过KeepAliveTimeout，关闭连接，我们设置为3s

PermitWithoutStream	//如果为true，当连接空闲时仍然发送PING帧监测，如果为false，则不发送忽略。我们设置为true
```

## 0x02 gRPC Pool 分析[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#0x02-grpc-pool-%E5%88%86%E6%9E%90)

滴滴开源的 [grpc 连接池](https://github.com/shimingyah/pool)，代码不长。简单分析下：

#### grpc.conn 封装[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#grpcconn-%E5%B0%81%E8%A3%85)

代码主要 [在此](https://github.com/shimingyah/pool/blob/master/conn.go)，

```
// Conn single grpc connection inerface
type Conn interface {
	// Value return the actual grpc connection type *grpc.ClientConn.
	Value() *grpc.ClientConn

	// Close decrease the reference of grpc connection, instead of close it.
	// if the pool is full, just close it.
	Close() error
}

// Conn is wrapped grpc.ClientConn. to provide close and value method.
type conn struct {
	cc   *grpc.ClientConn       // 封装真正的 grpc.conn
	pool *pool                  // 指向的 pool
	once bool
}
```

#### Pool 封装[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#pool-%E5%B0%81%E8%A3%85)

gRPC-Pool 封装的主要代码 [在此](https://github.com/shimingyah/pool/blob/master/pool.go)，一个 `interface`，一个 `struct`：  
Pool 对外部暴露的接口就 `3` 个：

- `Get`：从连接池获取连接
- `Close`：关闭连接池
- `Status`：打印连接池信息

```
// Pool interface describes a pool implementation.
// An ideal pool is threadsafe and easy to use.
type Pool interface {
	// Get returns a new connection from the pool. Closing the connections puts
	// it back to the Pool. Closing it when the pool is destroyed or full will
	// be counted as an error. we guarantee the conn.Value() isn't nil when conn isn't nil.
	Get() (Conn, error)     // 从池中取连接

	// Close closes the pool and all its connections. After Close() the pool is
	// no longer usable. You can't make concurrent calls Close and Get method.
	// It will be cause panic.
	Close() error           // 关闭池，关闭池中的连接

	// Status returns the current status of the pool.
	Status() string
}

// gRPC 连接池的定义
type pool struct {
	// atomic, used to get connection random
	index uint32
	// atomic, the current physical connection of pool
	current int32
	// atomic, the using logic connection of pool
	// logic connection = physical connection * MaxConcurrentStreams
	ref int32
	// pool options
	opt Options
    // all of created physical connections
    //  真正存储连接的结构
	conns []*conn
	// the server address is to create connection.
	address string
	// control the atomic var current's concurrent read write.
	sync.RWMutex
}
```

```
// New return a connection pool.
func New(address string, option Options) (Pool, error) {
	if address == "" {
		return nil, errors.New("invalid address settings")
	}
	if option.Dial == nil {
		return nil, errors.New("invalid dial settings")
	}
	if option.MaxIdle <= 0 || option.MaxActive <= 0 || option.MaxIdle> option.MaxActive {
		return nil, errors.New("invalid maximum settings")
	}
	if option.MaxConcurrentStreams <= 0 {
		return nil, errors.New("invalid maximun settings")
	}

	p := &pool{
		index:   0,
		current: int32(option.MaxIdle),
		ref:     0,
		opt:     option,
		conns:   make([]*conn, option.MaxActive),
		address: address,
	}

	for i := 0; i < p.opt.MaxIdle; i++ {
		c, err := p.opt.Dial(address)
		if err != nil {
			p.Close()
			return nil, fmt.Errorf("dial is not able to fill the pool: %s", err)
		}
		p.conns[i] = p.wrapConn(c, false)
	}
	log.Printf("new pool success: %v\n", p.Status())

	return p, nil
}

func (p *pool) incrRef() int32 {
	newRef := atomic.AddInt32(&p.ref, 1)
	if newRef == math.MaxInt32 {
		panic(fmt.Sprintf("overflow ref: %d", newRef))
	}
	return newRef
}

func (p *pool) decrRef() {
	newRef := atomic.AddInt32(&p.ref, -1)
	if newRef < 0 {
		panic(fmt.Sprintf("negative ref: %d", newRef))
	}
	if newRef == 0 && atomic.LoadInt32(&p.current) > int32(p.opt.MaxIdle) {
		p.Lock()
		if atomic.LoadInt32(&p.ref) == 0 {
			log.Printf("shrink pool: %d ---> %d, decrement: %d, maxActive: %d\n",
				p.current, p.opt.MaxIdle, p.current-int32(p.opt.MaxIdle), p.opt.MaxActive)
			atomic.StoreInt32(&p.current, int32(p.opt.MaxIdle))
			p.deleteFrom(p.opt.MaxIdle)
		}
		p.Unlock()
	}
}

func (p *pool) reset(index int) {
	conn := p.conns[index]
	if conn == nil {
		return
	}
	conn.reset()
	p.conns[index] = nil
}

func (p *pool) deleteFrom(begin int) {
	for i := begin; i < p.opt.MaxActive; i++ {
		p.reset(i)
	}
}

// Get see Pool interface.
func (p *pool) Get() (Conn, error) {
	// the first selected from the created connections
	nextRef := p.incrRef()
	p.RLock()
	current := atomic.LoadInt32(&p.current)
	p.RUnlock()
	if current == 0 {
		return nil, ErrClosed
	}
	if nextRef <= current*int32(p.opt.MaxConcurrentStreams) {
		next := atomic.AddUint32(&p.index, 1) % uint32(current)
		return p.conns[next], nil
	}

	// the number connection of pool is reach to max active
	if current == int32(p.opt.MaxActive) {
		// the second if reuse is true, select from pool's connections
		if p.opt.Reuse {
			next := atomic.AddUint32(&p.index, 1) % uint32(current)
			return p.conns[next], nil
		}
		// the third create one-time connection
		c, err := p.opt.Dial(p.address)
		return p.wrapConn(c, true), err
	}

	// the fourth create new connections given back to pool
	p.Lock()
	current = atomic.LoadInt32(&p.current)
	if current <int32(p.opt.MaxActive) && nextRef > current*int32(p.opt.MaxConcurrentStreams) {
		// 2 times the incremental or the remain incremental
		increment := current
		if current+increment > int32(p.opt.MaxActive) {
			increment = int32(p.opt.MaxActive) - current
		}
		var i int32
		var err error
		for i = 0; i < increment; i++ {
			c, er := p.opt.Dial(p.address)
			if er != nil {
				err = er
				break
			}
			p.reset(int(current + i))
			p.conns[current+i] = p.wrapConn(c, false)
		}
		current += i
		log.Printf("grow pool: %d ---> %d, increment: %d, maxActive: %d\n",
			p.current, current, increment, p.opt.MaxActive)
		atomic.StoreInt32(&p.current, current)
		if err != nil {
			p.Unlock()
			return nil, err
		}
	}
	p.Unlock()
	next := atomic.AddUint32(&p.index, 1) % uint32(current)
	return p.conns[next], nil
}

// Close see Pool interface.
func (p *pool) Close() error {
	atomic.StoreUint32(&p.index, 0)
	atomic.StoreInt32(&p.current, 0)
	atomic.StoreInt32(&p.ref, 0)
	p.deleteFrom(0)
	log.Printf("close pool success: %v\n", p.Status())
	return nil
}

// Status see Pool interface.
func (p *pool) Status() string {
	return fmt.Sprintf("address:%s, index:%d, current:%d, ref:%d. option:%v",
		p.address, p.index, p.current, p.ref, p.opt)
}
```

#### grpc.Pool 的使用[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#grpcpool-%E7%9A%84%E4%BD%BF%E7%94%A8)

本小节给出基于 gRPC 连接池的 CS 调用例子，如下：

服务端代码：

```
func main() {
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 调整 grpc 的默认参数
	s := grpc.NewServer(
		grpc.InitialWindowSize(pool.InitialWindowSize),
		grpc.InitialConnWindowSize(pool.InitialConnWindowSize),
		grpc.MaxSendMsgSize(pool.MaxSendMsgSize),
		grpc.MaxRecvMsgSize(pool.MaxRecvMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    pool.KeepAliveTime,
			Timeout: pool.KeepAliveTimeout,
		}),
	)
	pb.RegisterEchoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

客户端代码：

```
func main() {
	flag.Parse()

	p, err := pool.New(*addr, pool.DefaultOptions)
	if err != nil {
		log.Fatalf("failed to new pool: %v", err)
	}
	defer p.Close()

	conn, err := p.Get()
	if err != nil {
		log.Fatalf("failed to get conn: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn.Value())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.Say(ctx, &pb.EchoRequest{Message: []byte("hi")})
	if err != nil {
		log.Fatalf("unexpected error from Say: %v", err)
	}
	fmt.Println("rpc response:", res)
}
```

## 0x03 通用 TCP 连接池的实现[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#0x03-%E9%80%9A%E7%94%A8-tcp-%E8%BF%9E%E6%8E%A5%E6%B1%A0%E7%9A%84%E5%AE%9E%E7%8E%B0)

基于前面的分析，如何实现一个通用的 Tcp 连接池呢？以此项目[A golang general network connection poolction pool](https://github.com/silenceper/pool)

- 连接池中连接类型为`interface{}`，更通用
- 连接的最大空闲时间，超时的连接将关闭丢弃，可避免空闲时连接自动失效问题
- 支持用户设定 ping 方法，检查连接的连通性，无效的连接将丢弃
- 使用channel高效处理池中的连接

#### 使用方法[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#%E4%BD%BF%E7%94%A8%E6%96%B9%E6%B3%95)

```
//factory 创建连接的方法
factory := func() (interface{}, error) { 
	return net.Dial("tcp", "127.0.0.1:12345") 
}

//close 关闭连接的方法
close := func(v interface{}) error { 
	return v.(net.Conn).Close() 
}

//创建一个连接池： 初始化5，最大空闲连接是20，最大并发连接30
poolConfig := &pool.Config{
	InitialCap: 5,//资源池初始连接数
	MaxIdle:   20,//最大空闲连接数
	MaxCap:     30,//最大并发连接数
	Factory:    factory,
	Close:      close,
	//Ping:       ping,
	//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
	IdleTimeout: 15 * time.Second,
}
p, err := pool.NewChannelPool(poolConfig)
if err != nil {
	fmt.Println("err=", err)
}

//从连接池中取得一个连接
v, err := p.Get()

//do something
//conn=v.(net.Conn)

//将连接放回连接池中
p.Put(v)
//释放连接池中的所有连接
p.Release()
//查看当前连接中的数量
current := p.Len()
```

## 0x04 后记：是否需要gRPC的连接池[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#0x04-%E5%90%8E%E8%AE%B0%E6%98%AF%E5%90%A6%E9%9C%80%E8%A6%81grpc%E7%9A%84%E8%BF%9E%E6%8E%A5%E6%B1%A0)

由于现网中，笔者使用的场景大多数都是基于RPC的长连接（如Etcd/Consul/kubernetes-endpoint等），即gRPC 内建的 balancer 已经提供了优秀的连接管理支持（而且还可以自己实现池及Loadbalancer策略），每个后端实例一个 HTTP2 物理连接。**个人认为连接池机制比较适合于短连接的场景**

## 0x05 总结[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#0x05-%E6%80%BB%E7%BB%93)

- Redis/MySQL 连接池这种是单个连接只能负载一个并发，没有可用连接时会阻塞执行，并发跟不上的时候连接池相应调大点，性能会提升
- gRPC 内建的 balancer 已经有很好的连接管理的支持了，每个后端实例一个 HTTP2 物理连接
- gRPC 的 HTTP2 连接有复用能力，N 个 goroutine 用一个 HTTP2 连接没有任何问题，并不会单纯因为没有可用连接而阻塞

## 0x06 参考[](https://pandaychen.github.io/2020/10/03/DO-WE-NEED-GRPC-CLIENT-POOL/#0x06-%E5%8F%82%E8%80%83)

- [一个典型的 tcp 连接池开源实现](https://github.com/silenceper/pool)
- [gRPC 连接池](https://github.com/shimingyah/pool)
- [gRPC 连接池的设计与实现](https://zhuanlan.zhihu.com/p/100200985)
- [grpc-go-pool](https://github.com/processout/grpc-go-pool)
- [Golang 连接池的几种实现案例](https://zhuanlan.zhihu.com/p/109852936)
- [Connection pool for Go’s net.Conn interface](https://github.com/fatih/pool)
- [A golang general network connection poolction pool](https://github.com/silenceper/pool)