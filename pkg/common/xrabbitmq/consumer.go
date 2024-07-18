package xrabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
)

type RabbitConsumer struct {
	*Rabbitmq
	consumer Consumer
}

func NewRabbitConsumer(cfg *conf.Rabbitmq, consumer Consumer) (consumerClient *RabbitConsumer) {
	consumerClient = &RabbitConsumer{Rabbitmq: &Rabbitmq{cfg: cfg, waitingTime: 1}, consumer: consumer}
	consumerClient.connected = consumerClient.consuming
	err := consumerClient.connect()
	if err == nil {
		go consumerClient.notifyListen()
	}
	return consumerClient
}

func (r *RabbitConsumer) consuming() {
	var (
		queue    amqp.Queue
		delivery <-chan amqp.Delivery
		err      error
	)
	if queue, err = r.queueDeclare(r.cfg.Queue, nil); err != nil {
		xlog.Warnf("Failed to declare queue: %s", err.Error())
		return
	}
	if err = r.queueBind(queue.Name, r.cfg.RouteKey, r.cfg.Exchange); err != nil {
		xlog.Warnf("Failed to bind queue: %s", err.Error())
		return
	}
	if err = r.ch.Qos(1, 0, false); err != nil {
		xlog.Warnf("Failed to set qos: %s", err.Error())
		return
	}

	delivery, err = r.ch.Consume(
		queue.Name,        // name
		r.cfg.ConsumerTag, // consumerTag,
		false,             // noAck
		false,             // exclusive
		false,             // noLocal
		false,             // noWait
		nil,               // arguments
	)
	if err != nil {
		xlog.Warnf("Failed to consume: %s", err.Error())
		return
	}
	go func() {
		for {
			select {
			case d, ok := <-delivery:
				if ok == false {
					xlog.Warn(ERROR_AMQP_CONNECTION_WAS_LOST.Error())
					return
				}
				r.consumer(d)
				err = d.Ack(false)
				if err != nil {
					xlog.Warnf("Failed to ack: %s", err.Error())
				}
			}
		}
	}()
}
