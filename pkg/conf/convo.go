package conf

type Convo struct {
	ChanSize int  `yaml:"chan_size"`
	MapSize  int  `yaml:"map_size"`
	Interval int  `yaml:"interval"`
	Enabled  bool `yaml:"enabled"`
}
