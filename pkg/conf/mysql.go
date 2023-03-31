package conf

//type Mysql struct {
//	Address      string `yaml:"address"`
//	Username     string `yaml:"username"`
//	Password     string `yaml:"password"`
//	Db           string `yaml:"db"`
//	MaxOpenConn  int    `yaml:"max_open_conn"`
//	MaxIdleConn  int    `yaml:"max_idle_conn"`
//	ConnLifetime int    `yaml:"conn_lifetime"`
//	Charset      string `yaml:"charset"`
//}

type Mysql struct {
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxLifetime  int    `yaml:"max_lifetime"`
	MaxIdleTime  int    `yaml:"max_idle_time"`
	Charset      string `yaml:"charset"`
	Sources      []*Db  `yaml:"sources"`
	Replicas     []*Db  `yaml:"replicas"`
}

type Db struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}
