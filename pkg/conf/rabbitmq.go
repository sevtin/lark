package conf

type Rabbitmq struct {
	Address  []string `yaml:"address"`
	Password string   `yaml:"password"`
	Username string   `yaml:"username"`
	Vhost    string   `yaml:"vhost"`
	Exchange string   `yaml:"exchange"`  // 交换器名称
	RouteKey string   `yaml:"route_key"` // 路由名称
	Queue    string   `yaml:"queue"`     // 队列名称
}
