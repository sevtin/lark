package conf

type WsServer struct {
	ServerId              int    `yaml:"server_id" json:"server_id"`
	Name                  string `yaml:"name" json:"name"`
	Port                  int    `yaml:"port" json:"port"`
	WriteWait             int    `yaml:"write_wait" json:"write_wait"`
	PongWait              int    `yaml:"pong_wait" json:"pong_wait"`
	PingPeriod            int    `yaml:"ping_period" json:"ping_period"`
	MaxMessageSize        int    `yaml:"max_message_size" json:"max_message_size"`
	ReadBufferSize        int    `yaml:"read_buffer_size" json:"read_buffer_size"`
	WriteBufferSize       int    `yaml:"write_buffer_size" json:"write_buffer_size"`
	HeaderLength          int    `yaml:"header_length" json:"header_length"`
	ChanClientSendMessage int    `yaml:"chan_client_send_message" json:"chan_client_send_message"`
	ChanServerReadMessage int    `yaml:"chan_server_read_message" json:"chan_server_read_message"`
	ChanServerRegister    int    `yaml:"chan_server_register" json:"chan_server_register"`
	ChanServerUnregister  int    `yaml:"chan_server_unregister" json:"chan_server_unregister"`
	MaxConnections        int    `yaml:"max_connections" json:"max_connections"`
	MinimumTimeInterval   int    `yaml:"minimum_time_interval" json:"minimum_time_interval"`
	Log                   string `yaml:"log" json:"log"`
}
