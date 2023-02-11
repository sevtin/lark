package conf

type GrpcServer struct {
	Name string `yaml:"name"`
	Cert *Cert  `yaml:"cert"`
}
