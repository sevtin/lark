package conf

type AwsS3 struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	EndPoint  string `yaml:"endpoint"`
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket"`
}
