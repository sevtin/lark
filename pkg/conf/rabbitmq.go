package conf

type Rabbitmq struct {
	Address     []string `yaml:"address"`
	Username    string   `yaml:"username"`
	Password    string   `yaml:"password"`
	Vhost       string   `yaml:"vhost"`
	Exchange    string   `yaml:"exchange"`  // 交换器名称
	RouteKey    string   `yaml:"route_key"` // 路由名称
	Queue       string   `yaml:"queue"`     // 队列名称
	ConsumerTag string   `yaml:"consumer_tag"`
	Kind        string   `yaml:"kind"` // exchangeType
}
