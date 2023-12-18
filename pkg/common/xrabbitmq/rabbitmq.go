package xrabbitmq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var (
	ERROR_AMQP_CONN_IS_NIL         = errors.New("amqp connection is nil")
	ERROR_AMQP_CH_IS_NIL           = errors.New("amqp channel is nil")
	ERROR_AMQP_CONNECTION_WAS_LOST = errors.New("amqp connection was lost")
)

type RabbitClient struct {
	cfg        *conf.Rabbitmq
	conn       *amqp.Connection
	ch         *amqp.Channel
	notifyCh   chan *amqp.Error
	connecting bool
	consumer   Consumer
}

type Consumer func(amqp.Delivery)

func NewRabbitClient(cfg *conf.Rabbitmq) (client *RabbitClient) {
	client = &RabbitClient{cfg: cfg}
	err := client.connect()
	if err == nil {
		go client.notifyListen()
	}
	return client
}

func (mq *RabbitClient) connect() (err error) {
	var (
		conn *amqp.Connection
		ch   *amqp.Channel
	)
	if mq.connecting == true {
		return
	}
	mq.connecting = true
	defer func() {
		mq.connecting = false
	}()
	mq.close()

	if conn, err = amqp.Dial(getAMQPUrl(mq.cfg)); err != nil {
		xlog.Warn(err.Error())
		return
	}
	if ch, err = conn.Channel(); err != nil {
		xlog.Warn(err.Error())
		return
	}
	mq.conn = conn
	mq.ch = ch
	mq.notifyCh = conn.NotifyClose(make(chan *amqp.Error))
	if err = mq.declareExchange(); err != nil {
		xlog.Warn(err.Error())
		return
	}
	if mq.consumer != nil {
		mq.consuming()
	}
	return
}

func (mq *RabbitClient) notifyListen() {
	var (
		err *amqp.Error
	)

	defer func() {
		if r := recover(); r != nil {
			xlog.Warn(string(debug.Stack()))
		}
	}()

	for {
		select {
		case err = <-mq.notifyCh:
			if err != nil {
				xlog.Warn(err.Error())
			}
			xlog.Warn("Reconnecting to RabbitMQ...")
			mq.connect()
			time.Sleep(2 * time.Second)
		}
	}
}

func getAMQPUrl(cfg *conf.Rabbitmq) (url string) {
	url = fmt.Sprintf(
		"amqp://%s:%s@%s/%s",
		cfg.Username,
		cfg.Password,
		strings.Join(cfg.Address, ","),
		cfg.Vhost,
	)
	return
}

func (mq *RabbitClient) Send(msg []byte) (err error) {
	return mq.publishMessage(mq.cfg.Exchange, mq.cfg.RouteKey, msg)
}

func (mq *RabbitClient) SendDelay(msg []byte, delayTime int) (err error) {
	var (
		queue          amqp.Queue
		expiration     = strconv.Itoa(delayTime)
		delayQueueName = mq.cfg.Queue + "_delay:" + expiration
		delayRouteKey  = mq.cfg.RouteKey + "_delay:" + expiration
		args           = amqp.Table{
			"x-dead-letter-exchange":    mq.cfg.Exchange,
			"x-dead-letter-routing-key": mq.cfg.RouteKey,
		}
	)
	if queue, err = mq.declareQueue(delayQueueName, args); err != nil {
		return
	}
	if err = mq.bindQueue(queue.Name, delayRouteKey, mq.cfg.Exchange); err != nil {
		return
	}
	return mq.publishMessage(mq.cfg.Exchange, delayRouteKey, msg, amqp.Publishing{Expiration: expiration})
}

func (mq *RabbitClient) publishMessage(exchange, routeKey string, body []byte, options ...amqp.Publishing) (err error) {
	if mq.ch == nil {
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
	return mq.ch.Publish(exchange, routeKey, false, false, publishing)
}

func (mq *RabbitClient) Consume(consumer Consumer) {
	mq.consumer = consumer
	mq.consuming()
}

func (mq *RabbitClient) consuming() {
	var (
		queue    amqp.Queue
		delivery <-chan amqp.Delivery
		err      error
	)
	if queue, err = mq.declareQueue(mq.cfg.Queue, nil); err != nil {
		xlog.Warn(err.Error())
		return
	}
	if err = mq.bindQueue(queue.Name, mq.cfg.RouteKey, mq.cfg.Exchange); err != nil {
		xlog.Warn(err.Error())
		return
	}
	if err = mq.ch.Qos(1, 0, false); err != nil {
		xlog.Warn(err.Error())
		return
	}

	delivery, err = mq.ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		xlog.Warn(err.Error())
		return
	}
	go func() {
		var (
			d  amqp.Delivery
			ok bool
		)
		for {
			select {
			case d, ok = <-delivery:
				if ok == false {
					xlog.Warn(ERROR_AMQP_CONNECTION_WAS_LOST.Error())
					return
				}
				mq.consumer(d)
				err = d.Ack(false)
				if err != nil {
					xlog.Warn(err.Error())
				}
			}
		}
	}()
}

func (mq *RabbitClient) close() {
	if mq.ch != nil {
		mq.ch.Close()
	}
	if mq.conn != nil {
		mq.conn.Close()
	}
}

func (mq *RabbitClient) declareQueue(name string, args amqp.Table) (queue amqp.Queue, err error) {
	if mq.ch == nil {
		err = ERROR_AMQP_CH_IS_NIL
		return
	}
	queue, err = mq.ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		args,
	)
	return
}

func (mq *RabbitClient) declareExchange() (err error) {
	if mq.ch == nil {
		err = ERROR_AMQP_CH_IS_NIL
		return
	}
	err = mq.ch.ExchangeDeclare(
		mq.cfg.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	return
}

func (mq *RabbitClient) bindQueue(queue, routeKey, exchange string) (err error) {
	if mq.ch == nil {
		err = ERROR_AMQP_CH_IS_NIL
		return
	}
	err = mq.ch.QueueBind(
		queue,
		routeKey,
		exchange,
		false,
		nil,
	)
	return
}
