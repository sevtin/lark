package gateway

import (
	"github.com/Shopify/sarama"
	"google.golang.org/grpc"
	"io"
	"lark/apps/msg_gateway/internal/config"
	"lark/apps/msg_gateway/internal/server/websocket/ws"
	"lark/apps/msg_gateway/internal/service"
	"lark/domain/cache"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xkafka"
	"lark/pkg/common/xmonitor"
	"lark/pkg/constant"
	"lark/pkg/obj"
	"lark/pkg/proto/pb_gw"
	"lark/pkg/utils"
	"time"
)

type GatewayServer interface {
	Run()
}

type gatewayServer struct {
	pb_gw.UnimplementedMessageGatewayServer
	conf           *config.Config
	wsServer       *ws.WServer
	wsService      service.WsService
	grpcServer     *xgrpc.GrpcServer
	producer       *xkafka.Producer
	msgHandle      map[string]obj.KafkaMessageHandler
	consumerGroup  *xkafka.MConsumerGroup
	serverMgrCache cache.ServerMgrCache
}

func NewGatewayServer(conf *config.Config, wsService service.WsService, serverMgrCache cache.ServerMgrCache) GatewayServer {
	srv := &gatewayServer{conf: conf, wsService: wsService, serverMgrCache: serverMgrCache}
	conf.WsServer.ServerId = conf.ServerID
	conf.WsServer.Log = conf.Log
	srv.wsServer = ws.NewWServer(conf.WsServer, wsService.MessageCallback)
	srv.producer = xkafka.NewKafkaProducer(conf.MsgProducer.Address, conf.MsgProducer.Topic)
	srv.msgHandle = make(map[string]obj.KafkaMessageHandler)

	var (
		topics = make([]string, len(conf.MsgConsumer.Topic))
		topic  string
		i      int
		idStr  = utils.IntToStr(conf.GrpcServer.ServerID)
	)
	for i, topic = range conf.MsgConsumer.Topic {
		topic = topic + "_" + idStr
		srv.msgHandle[topic] = srv.MessageHandler
		topics[i] = topic
	}

	srv.consumerGroup = xkafka.NewMConsumerGroup(&xkafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_2_1_0, OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		topics,
		conf.MsgConsumer.Address,
		conf.MsgConsumer.GroupID)
	srv.consumerGroup.RegisterHandler(srv)

	return srv
}

func (s *gatewayServer) Setup(_ sarama.ConsumerGroupSession) error {
	close(s.consumerGroup.Ready)
	return nil
}
func (s *gatewayServer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (s *gatewayServer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		msg *sarama.ConsumerMessage
		err error
	)
	for {
		select {
		case msg = <-claim.Messages():
			if msg == nil {
				continue
			}
			if err = s.msgHandle[msg.Topic](msg.Value, string(msg.Key)); err != nil {
				continue
			}
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
	return nil
}

func (s *gatewayServer) Run() {
	xmonitor.RunMonitor(s.conf.Monitor.Port)
	s.conf.GrpcServer.Name = s.conf.GrpcServer.Name + ":" + utils.IntToStr(s.conf.GrpcServer.ServerID)

	s.timedTask()
	s.wsServer.Run()
	s.wsService.Run()
	var (
		srv    *grpc.Server
		closer io.Closer
	)

	srv, closer = xgrpc.NewServer(s.conf.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()

	pb_gw.RegisterMessageGatewayServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.conf.GrpcServer, s.conf.Etcd)
	s.grpcServer.RunServer(srv)
}

func (s *gatewayServer) timedTask() {
	go func() {
		var (
			member = s.conf.GrpcServer.Name + ":" + utils.IntToStr(s.conf.WsServer.Port)
			num    int
		)
		s.serverMgrCache.ZAddMsgGateway(0, member)
		onlineTicker := time.NewTicker(constant.CONST_DURATION_NUMBER_OF_CHAT_MEMBER_ONLINE_SECOND)
		for {
			select {
			case <-onlineTicker.C:
				num = s.wsServer.NumberOfOnline()
				s.serverMgrCache.ZAddMsgGateway(float64(num), member)
			}
		}
	}()
}
