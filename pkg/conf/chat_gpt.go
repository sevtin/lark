package conf

type ChatGpt struct {
	OpenaiKeys []string `yaml:"openai_keys"`
	ApiUrl     string   `yaml:"api_url"`
	HttpProxy  string   `yaml:"http_proxy"`
}
