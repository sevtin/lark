package conf

type Etcd struct {
	Endpoints    []string `yaml:"endpoints"`
	Username     string   `yaml:"username"`
	Password     string   `yaml:"password"`
	Schema       string   `yaml:"schema"`
	ReadTimeout  int      `yaml:"read_timeout"`
	WriteTimeout int      `yaml:"write_timeout"`
	DialTimeout  int      `yaml:"dial_timeout"`
}
