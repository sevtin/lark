package conf

type Pg struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Dbname       string `yaml:"dbname"`
	Sslmode      string `yaml:"sslmode"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxLifetime  int    `yaml:"max_lifetime"`
	MaxIdleTime  int    `yaml:"max_idle_time"`
	Charset      string `yaml:"charset"`
	LogLevel     int    `yaml:"log_level"`
}
