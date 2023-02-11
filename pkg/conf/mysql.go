package conf

type Mysql struct {
	Address      string `yaml:"address"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Db           string `yaml:"db"`
	MaxOpenConn  int    `yaml:"max_open_conn"`
	MaxIdleConn  int    `yaml:"max_idle_conn"`
	ConnLifetime int    `yaml:"conn_lifetime"`
	Charset      string `yaml:"charset"`
}
