package conf

type Credential struct {
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
	Enabled  bool   `yaml:"enabled"`
}

type Cert struct {
	CertFile           string `yaml:"cert_file"`
	ServerNameOverride string `yaml:"server_name_override"`
	Enabled            bool   `yaml:"enabled"`
}
