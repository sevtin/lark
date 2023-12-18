package config

type FILE_TYPE int8

const (
	FILE_TYPE_GO    FILE_TYPE = 0
	FILE_TYPE_PROTO FILE_TYPE = 1
	FILE_TYPE_YAML  FILE_TYPE = 2
)

type GenConfig struct {
	Path        string // 文件保存路径
	Prefix      string // 前缀
	Suffix      string // 后缀
	PackageName string // 包名
	ServiceName string // 服务名称
	ApiName     string // api名称
	ModelName   string // 表模型
	Filename    string // 文件名
	FileType    FILE_TYPE
	Dict        map[string]interface{}
}

func (c *GenConfig) ResetPrefSuf() {
	c.Prefix = ""
	c.Suffix = ""
}
