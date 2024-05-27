package conf

type Mysql struct {
	Address      string `yaml:"address"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Db           string `yaml:"db"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxLifetime  int    `yaml:"max_lifetime"`
	MaxIdleTime  int    `yaml:"max_idle_time"`
	Charset      string `yaml:"charset"`
	LogLevel     int    `yaml:"log_level"`
}

//type Mysql struct {
//	MaxOpenConns int    `yaml:"max_open_conns"`
//	MaxIdleConns int    `yaml:"max_idle_conns"`
//	MaxLifetime  int    `yaml:"max_lifetime"`
//	MaxIdleTime  int    `yaml:"max_idle_time"`
//	Charset      string `yaml:"charset"`
//	Sources      []*Db  `yaml:"sources"`
//	Replicas     []*Db  `yaml:"replicas"`
//}
//
//type Db struct {
//	Address  string `yaml:"address"`
//	Username string `yaml:"username"`
//	Password string `yaml:"password"`
//	Db       string `yaml:"db"`
//}
