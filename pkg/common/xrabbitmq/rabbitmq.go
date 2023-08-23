package xrabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"lark/pkg/conf"
	"log"
	"strconv"
	"strings"
)

type RabbitClient struct {
	cfg  *conf.Rabbitmq
	conn *amqp.Connection
	ch   *amqp.Channel
}

type Consumer func(amqp.Delivery)

func NewRabbitClient(cfg *conf.Rabbitmq) (client *RabbitClient) {
	var (
		conn *amqp.Connection
		ch   *amqp.Channel
		err  error
	)
	if conn, err = amqp.Dial(getAMQPUrl(cfg)); err != nil {
		log.Panic(err)
		return
	}
	if ch, err = conn.Channel(); err != nil {
		log.Panic(err)
		return
	}
	client = &RabbitClient{cfg: cfg, conn: conn, ch: ch}
	if err = client.declareExchange(); err != nil {
		log.Panic(err)
		return
	}
	return client
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
	publishing := amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}
	if len(options) > 0 {
		publishing.Expiration = options[0].Expiration
	}
	return mq.ch.Publish(exchange, routeKey, false, false, publishing)
}

func (mq *RabbitClient) Consume(callback Consumer) {
	var (
		queue amqp.Queue
		msgs  <-chan amqp.Delivery
		err   error
	)
	if queue, err = mq.declareQueue(mq.cfg.Queue, nil); err != nil {
		log.Panic(err)
		return
	}
	if err = mq.bindQueue(queue.Name, mq.cfg.RouteKey, mq.cfg.Exchange); err != nil {
		log.Panic(err)
		return
	}
	if err = mq.ch.Qos(1, 0, false); err != nil {
		log.Panic(err)
		return
	}

	msgs, err = mq.ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
		return
	}
	go func() {
		for d := range msgs {
			callback(d)
			d.Ack(false)
		}
	}()
}

func (mq *RabbitClient) Close() {
	mq.ch.Close()
	mq.conn.Close()
}

func (mq *RabbitClient) declareQueue(name string, args amqp.Table) (queue amqp.Queue, err error) {
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
	err = mq.ch.QueueBind(
		queue,
		routeKey,
		exchange,
		false,
		nil,
	)
	return
}
