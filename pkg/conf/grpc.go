package conf

type Grpc struct {
	Name                  string      `yaml:"name"`
	ServerID              int         `yaml:"server_id"`
	Port                  int         `yaml:"port"`
	MaxConnectionIdle     int         `yaml:"max_connection_idle"`      // The current default value is infinity.
	MaxConnectionAge      int         `yaml:"max_connection_age"`       // The current default value is infinity.
	MaxConnectionAgeGrace int         `yaml:"max_connection_age_grace"` // The current default value is infinity.
	Time                  int         `yaml:"time"`                     // The current default value is 2 hours.
	Timeout               int         `yaml:"timeout"`                  // The current default value is 20 seconds.
	ConnectionLimit       int         `yaml:"connection_limit"`
	StreamsLimit          uint32      `yaml:"streams_limit"`
	MaxRecvMsgSize        int         `yaml:"max_recv_msg_size"`
	Credential            *Credential `yaml:"credential"`
	Jaeger                *Jaeger     `yaml:"jaeger"`
}
