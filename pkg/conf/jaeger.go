package conf

type Jaeger struct {
	HostPort            string  `yaml:"host_port"`
	SamplerType         string  `yaml:"sampler_type"`
	Param               float64 `yaml:"param"`
	LogSpans            bool    `yaml:"log_spans"`
	BufferFlushInterval int     `yaml:"buffer_flush_interval"`
	MaxPacketSize       int     `yaml:"max_packet_size"`
	Enabled             bool    `yaml:"enabled"`
}
