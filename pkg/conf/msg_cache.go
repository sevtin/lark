package conf

type MsgCache struct {
	L1Duration int64 `yaml:"l1_duration"` // redis缓存时长(毫秒)
	L2Duration int64 `yaml:"l2_duration"` // mongo热数据时长(毫秒)
}
