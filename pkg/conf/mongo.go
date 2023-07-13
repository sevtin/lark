package conf

type Mongo struct {
	Hosts             []string `yaml:"hosts"`
	Username          string   `yaml:"username"`
	Password          string   `yaml:"password"`
	ReplicaSet        string   `yaml:"replica_set"`
	Db                string   `yaml:"db"`
	Direct            bool     `yaml:"direct"`
	Timeout           int      `yaml:"timeout"`
	MaxPoolSize       int      `yaml:"max_pool_size"`
	RetainChatRecords int      `yaml:"retain_chat_records"`
	AuthSource        string   `yaml:"auth_source"`
}
