package xkafka

import (
	"context"
	"github.com/Shopify/sarama"
	"lark/pkg/common/xlog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	BalanceStrategySticky     = "sticky"
	BalanceStrategyRoundrobin = "roundrobin"
	BalanceStrategyRange      = "range"
)

type MConsumerGroup struct {
	sarama.ConsumerGroup
	groupID string
	topics  []string
	Ready   chan bool
}

type MConsumerGroupConfig struct {
	KafkaVersion    sarama.KafkaVersion
	OffsetsInitial  int64
	IsReturnErr     bool
	Assignor        string
	BalanceStrategy string
}

func NewMConsumerGroup(conf *MConsumerGroupConfig, topics, addrs []string, groupID string) (group *MConsumerGroup) {
	group = &MConsumerGroup{
		ConsumerGroup: nil,
		groupID:       groupID,
		topics:        topics,
		Ready:         make(chan bool),
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
	config.Version = conf.KafkaVersion
	config.Consumer.Offsets.Initial = conf.OffsetsInitial
	config.Consumer.Return.Errors = conf.IsReturnErr
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	switch conf.Assignor {
	case BalanceStrategySticky:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case BalanceStrategyRoundrobin:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case BalanceStrategyRange:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	}
	consumerGroup, err = sarama.NewConsumerGroup(addrs, groupID, config)
	if err != nil {
		xlog.Errorf("Error creating consumer group client: %v", err)
		return
	}
	group.ConsumerGroup = consumerGroup
	return
}

func (mc *MConsumerGroup) RegisterHandler(handler sarama.ConsumerGroupHandler) {
	go mc.RegisterConsumerGroupHandler(handler)
}

func (mc *MConsumerGroup) RegisterConsumerGroupHandler(handler sarama.ConsumerGroupHandler) {
	var (
		ctx, cancel         = context.WithCancel(context.Background())
		consumptionIsPaused = false
		wg                  = &sync.WaitGroup{}
		keepRunning         = true
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		for {
			if mc.ConsumerGroup == nil {
				xlog.Error("consumer group is null")
				break
			}
			if err = mc.ConsumerGroup.Consume(ctx, mc.topics, handler); err != nil {
				xlog.Errorf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			mc.Ready = make(chan bool)
		}
	}()

	// Await till the consumer has been set up
	<-mc.Ready
	// Sarama consumer up and running!...

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

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
		case <-sigusr1:
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

func (mc *MConsumerGroup) Close() {
	if err := mc.ConsumerGroup.Close(); err != nil {
		xlog.Errorf("Error closing client: %v", err)
	}
}
