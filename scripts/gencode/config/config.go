package config

type GenConfig struct {
	Path        string // 文件保存路径
	Prefix      string // 前缀
	Suffix      string // 后缀
	PackageName string // 包名
	FileType    int
	Dict        map[string]interface{}
}

const (
	FILE_TYPE_GO = iota
	FILE_TYPE_PROTO
	FILE_TYPE_YAML
)
