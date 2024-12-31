```
func NewConfig() *Config {
	c := &Config{}

	c.Admin.Retry.Max = 5
	c.Admin.Retry.Backoff = 100 * time.Millisecond
	c.Admin.Timeout = 3 * time.Second

	c.Net.MaxOpenRequests = 5
	c.Net.DialTimeout = 30 * time.Second
	c.Net.ReadTimeout = 30 * time.Second
	c.Net.WriteTimeout = 30 * time.Second
	c.Net.SASL.Handshake = true
	c.Net.SASL.Version = SASLHandshakeV1

	c.Metadata.Retry.Max = 3
	c.Metadata.Retry.Backoff = 250 * time.Millisecond
	c.Metadata.RefreshFrequency = 10 * time.Minute
	c.Metadata.Full = true
	c.Metadata.AllowAutoTopicCreation = true

	c.Producer.MaxMessageBytes = 1024 * 1024
	c.Producer.RequiredAcks = WaitForLocal
	c.Producer.Timeout = 10 * time.Second
	c.Producer.Partitioner = NewHashPartitioner
	c.Producer.Retry.Max = 3
	c.Producer.Retry.Backoff = 100 * time.Millisecond
	c.Producer.Return.Errors = true
	c.Producer.CompressionLevel = CompressionLevelDefault

	c.Producer.Transaction.Timeout = 1 * time.Minute
	c.Producer.Transaction.Retry.Max = 50
	c.Producer.Transaction.Retry.Backoff = 100 * time.Millisecond

	c.Consumer.Fetch.Min = 1
	c.Consumer.Fetch.Default = 1024 * 1024
	c.Consumer.Retry.Backoff = 2 * time.Second
	c.Consumer.MaxWaitTime = 500 * time.Millisecond
	c.Consumer.MaxProcessingTime = 100 * time.Millisecond
	c.Consumer.Return.Errors = false
	c.Consumer.Offsets.AutoCommit.Enable = true
	c.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	c.Consumer.Offsets.Initial = OffsetNewest
	c.Consumer.Offsets.Retry.Max = 3

	c.Consumer.Group.Session.Timeout = 10 * time.Second
	c.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	c.Consumer.Group.Rebalance.GroupStrategies = []BalanceStrategy{NewBalanceStrategyRange()}
	c.Consumer.Group.Rebalance.Timeout = 60 * time.Second
	c.Consumer.Group.Rebalance.Retry.Max = 4
	c.Consumer.Group.Rebalance.Retry.Backoff = 2 * time.Second
	c.Consumer.Group.ResetInvalidOffsets = true

	c.ClientID = defaultClientID
	c.ChannelBufferSize = 256
	c.ApiVersionsRequest = true
	c.Version = DefaultVersion
	c.MetricRegistry = metrics.NewRegistry()

	return c
}
```