package xkafka

import (
	"context"
	"errors"
	"github.com/IBM/sarama"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
)

const (
	BalanceStrategySticky     = "sticky"
	BalanceStrategyRoundrobin = "roundrobin"
	BalanceStrategyRange      = "range"
)

type ConsumerMessageHandler func(topic string, msg *sarama.ConsumerMessage) (err error)

type ConsumerGroup struct {
	sarama.ConsumerGroup
	groupID                string
	topics                 []string
	ready                  chan bool
	consumerMessageHandler ConsumerMessageHandler
}

func NewConsumerGroup(conf *conf.KafkaConsumer, consumerMessageHandler ConsumerMessageHandler) (group *ConsumerGroup) {
	group = &ConsumerGroup{
		ConsumerGroup:          nil,
		groupID:                conf.GroupID,
		topics:                 conf.Topic,
		ready:                  make(chan bool),
		consumerMessageHandler: consumerMessageHandler,
	}
	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	var (
		config        = sarama.NewConfig()
		consumerGroup sarama.ConsumerGroup
		err           error
	)
	/*
		使用BalanceStrategySticky策略，这有助于在消费者群组成员变动时尽量保持分区分配不变，减少再平衡的频率。
	*/
	if conf.Sasl != nil && conf.Sasl.Enable == true {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = conf.Sasl.User
		config.Net.SASL.Password = conf.Sasl.Password
	}
	config.Version = conf.KafkaVersion
	config.Consumer.Offsets.Initial = conf.OffsetsInitial
	config.Consumer.Return.Errors = conf.IsReturnErr
	// config.Consumer.Offsets.AutoCommit.Enable = false // 关闭自动提交偏移量，以便手动控制
	switch conf.Assignor {
	case BalanceStrategySticky:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
		//config.Producer.Partitioner = sarama.NewManualPartitioner
	case BalanceStrategyRoundrobin:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
		//config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	case BalanceStrategyRange:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
		//config.Producer.Partitioner = sarama.NewManualPartitioner
	default:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
		//config.Producer.Partitioner = sarama.NewManualPartitioner
	}

	consumerGroup, err = sarama.NewConsumerGroup(conf.Address, conf.GroupID, config)
	if err != nil {
		xlog.Errorf("Error creating consumer group client: %v", err)
		return
	}
	group.ConsumerGroup = consumerGroup
	group.registerHandler()
	return
}

func (mc *ConsumerGroup) registerHandler() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				xlog.Errorf("Panic in goroutine: %v; Stack trace: %s", r, debug.Stack())
			}
		}()
		mc.registerConsumerGroupHandler()
	}()
}

func (mc *ConsumerGroup) registerConsumerGroupHandler() {
	var (
		ctx, cancel         = context.WithCancel(context.Background())
		consumptionIsPaused = false
		wg                  = &sync.WaitGroup{}
		keepRunning         = true
	)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				xlog.Errorf("Panic in goroutine: %v; Stack trace: %s", r, debug.Stack())
			}
		}()
		var err error
		for {
			if mc.ConsumerGroup == nil {
				xlog.Error("consumer group is null")
				break
			}
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err = mc.ConsumerGroup.Consume(ctx, mc.topics, mc); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				xlog.Errorf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			mc.ready = make(chan bool)
		}
	}()

	// Await till the consumer has been set up
	<-mc.ready
	// Sarama consumer up and running!...

	sigusr := make(chan os.Signal, 1)
	signal.Notify(sigusr, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			// terminating: context cancelled
			keepRunning = false
		case <-sigterm:
			// terminating: via signal
			keepRunning = false
		case <-sigusr:
			toggleConsumptionFlow(mc.ConsumerGroup, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()
	mc.Close()
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		// Resuming consumption
	} else {
		client.PauseAll()
		// Pausing consumption
	}
	*isPaused = !*isPaused
}

func (mc *ConsumerGroup) Close() {
	if err := mc.ConsumerGroup.Close(); err != nil {
		xlog.Errorf("Error closing client: %v", err)
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (mc *ConsumerGroup) Setup(sarama.ConsumerGroupSession) error {
	close(mc.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (mc *ConsumerGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (mc *ConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if ok == false {
				xlog.Info("message channel was closed")
				return nil
			}
			mc.messageHandler(claim.Topic(), msg)
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
	return nil
}

func (mc *ConsumerGroup) messageHandler(topic string, msg *sarama.ConsumerMessage) (err error) {
	defer func() {
		if r := recover(); r != nil {
			xlog.Errorf("Panic in goroutine: %v; Stack trace: %s", r, debug.Stack())
		}
	}()
	return mc.consumerMessageHandler(topic, msg)
}
