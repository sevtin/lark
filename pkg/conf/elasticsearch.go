package conf

type Elasticsearch struct {
	Addresses  []string `yaml:"address"`
	Username   string   `yaml:"username"`
	Password   string   `yaml:"password"`
	CACert     string   `yaml:"ca_cert"`
	TlsEnabled bool     `yaml:"tls_enabled"`
	//MaxIdleConnsPerHost   int      `yaml:"max_idle_conns_per_host"`
	//ResponseHeaderTimeout int      `yaml:"response_header_timeout"`
	//MaxVersion            int      `yaml:"max_version"`
	//InsecureSkipVerify    bool     `yaml:"insecure_skip_verify"`
}
