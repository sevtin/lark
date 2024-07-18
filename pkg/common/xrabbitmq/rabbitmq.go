package xrabbitmq

import (
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"runtime/debug"
	"strings"
	"time"
)

var (
	ERROR_AMQP_CONN_IS_NIL         = errors.New("amqp connection is nil")
	ERROR_AMQP_CH_IS_NIL           = errors.New("amqp channel is nil")
	ERROR_AMQP_CONNECTION_WAS_LOST = errors.New("amqp connection was lost")
)

const MaxWaitingTime = 64 // 秒

type Consumer func(amqp.Delivery)
type SuccessfullyConnected func()

type Rabbitmq struct {
	cfg         *conf.Rabbitmq
	conn        *amqp.Connection
	ch          *amqp.Channel
	notifyCh    chan *amqp.Error
	connecting  bool
	connected   SuccessfullyConnected
	waitingTime time.Duration
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

func (r *Rabbitmq) connect() (err error) {
	var (
		conn *amqp.Connection
		ch   *amqp.Channel
	)
	if r.cfg.Kind == "" {
		r.cfg.Kind = "direct"
	}
	if r.connecting == true {
		return
	}
	r.connecting = true
	defer func() {
		r.connecting = false
	}()
	r.close()

	if conn, err = amqp.Dial(getAMQPUrl(r.cfg)); err != nil {
		xlog.Warnf("amqp dial error: %s", err.Error())
		return
	}
	if ch, err = conn.Channel(); err != nil {
		xlog.Warnf("amqp channel error: %s", err.Error())
		return
	}
	r.conn = conn
	r.ch = ch
	r.notifyCh = conn.NotifyClose(make(chan *amqp.Error))
	if err = r.exchangeDeclare(); err != nil {
		xlog.Errorf("amqp declare exchange error: %s", err.Error())
		return
	}
	r.waitingTime = 1
	if r.connected == nil {
		return
	}
	r.connected()
	return
}

func (r *Rabbitmq) notifyListen() {
	defer func() {
		if re := recover(); re != nil {
			xlog.Warn(debug.Stack())
		}
	}()

	for {
		select {
		case err := <-r.notifyCh:
			if err != nil {
				xlog.Warnf("amqp notify error: %s", err.Error())
			}
			r.waitingTime *= 2
			if r.waitingTime > MaxWaitingTime {
				r.waitingTime = MaxWaitingTime
			}
			time.Sleep(r.waitingTime * time.Second)
			xlog.Warn("Reconnecting to RabbitMQ...")
			r.connect()
		}
	}
}

func (r *Rabbitmq) close() {
	if r.ch != nil {
		r.ch.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

func (r *Rabbitmq) exchangeDeclare() (err error) {
	if r.ch == nil {
		err = ERROR_AMQP_CH_IS_NIL
		return
	}
	err = r.ch.ExchangeDeclare(
		r.cfg.Exchange, // name
		r.cfg.Kind,     // exchangeType "direct", "Exchange type - direct|fanout|topic|x-custom"
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // noWait
		nil,            // arguments
	)
	return
}

func (r *Rabbitmq) queueDeclare(name string, args amqp.Table) (queue amqp.Queue, err error) {
	if r.ch == nil {
		err = ERROR_AMQP_CH_IS_NIL
		return
	}
	queue, err = r.ch.QueueDeclare(
		name,  // name of the queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		args,  // arguments
	)
	return
}

func (r *Rabbitmq) queueBind(queue, routeKey, exchange string) (err error) {
	if r.ch == nil {
		err = ERROR_AMQP_CH_IS_NIL
		return
	}
	err = r.ch.QueueBind(
		queue,    // name of the queue
		routeKey, // bindingKey
		exchange, // sourceExchange
		false,    // noWait
		nil,      // arguments
	)
	return
}
