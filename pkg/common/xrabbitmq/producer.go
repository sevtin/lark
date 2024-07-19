package xrabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"strconv"
)

type RabbitProducer struct {
	*Rabbitmq
}

func NewRabbitProducer(cfg *conf.Rabbitmq) (producer *RabbitProducer) {
	producer = &RabbitProducer{&Rabbitmq{cfg: cfg, waitingTime: 1}}
	err := producer.connect()
	if err == nil {
		go producer.notifyListen()
	}
	return producer
}

func (r *RabbitProducer) Send(msg []byte) (err error) {
	if len(msg) == 0 {
		return
	}
	return r.publish(r.cfg.Exchange, r.cfg.RouteKey, msg)
}

func (r *RabbitProducer) SendDelay(msg []byte, delayTime int64) (err error) {
	if len(msg) == 0 {
		return
	}
	var (
		queue          amqp.Queue
		expiration     = strconv.FormatInt(delayTime, 10)
		delayQueueName = r.cfg.Queue + "_delay"
		delayRouteKey  = r.cfg.RouteKey + "_delay"
		args           = amqp.Table{
			"x-dead-letter-exchange":    r.cfg.Exchange,
			"x-dead-letter-routing-key": r.cfg.RouteKey,
		}
	)
	if queue, err = r.queueDeclare(delayQueueName, args); err != nil {
		xlog.Warnf("Failed to declare delay queue: %s", err.Error())
		return
	}
	if err = r.queueBind(queue.Name, delayRouteKey, r.cfg.Exchange); err != nil {
		xlog.Warnf("Failed to bind delay queue: %s", err.Error())
		return
	}
	return r.publish(r.cfg.Exchange, delayRouteKey, msg, amqp.Publishing{Expiration: expiration})
}

func (r *RabbitProducer) publish(exchange, routeKey string, body []byte, options ...amqp.Publishing) (err error) {
	if r.ch == nil {
		err = ERROR_AMQP_CH_IS_NIL
		return
	}
	publishing := amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}
	if len(options) > 0 {
		publishing.Expiration = options[0].Expiration
	}
	return r.ch.Publish(exchange, routeKey, false, false, publishing)
}
