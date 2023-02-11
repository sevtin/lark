package conf

type Mongo struct {
	Address           string `yaml:"address"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Db                string `yaml:"db"`
	Direct            bool   `yaml:"direct"`
	Timeout           int    `yaml:"timeout"`
	MaxPoolSize       int    `yaml:"max_pool_size"`
	RetainChatRecords int    `yaml:"retain_chat_records"`
}
