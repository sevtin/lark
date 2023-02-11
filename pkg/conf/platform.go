package conf

type Platform struct {
	Type        int    `yaml:"type"` // 1:Android 2:iOS
	Name        string `yaml:"name"`
	OfflinePush bool   `yaml:"offline_push"`
}
