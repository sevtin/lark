package conf

type KafkaProducer struct {
	Address []string `yaml:"address"`
	Topic   string   `yaml:"topic"` //生成 主题
}

type KafkaConsumer struct {
	Address []string `yaml:"address"`
	Topic   []string `yaml:"topic"` //消费 主题
	GroupID string   `yaml:"group_id"`
}
