package conf

type Pay struct {
	Alipay *Alipay `yaml:"alipay"`
	Paypal *Paypal `yaml:"paypal"`
}

type Alipay struct {
	Appid               string `yaml:"appid"`
	AppPrivateKey       string `yaml:"app_private_key"`
	AppCertPublicKey    string `yaml:"app_cert_public_key"`
	AlipayRootCert      string `yaml:"alipay_root_cert"`
	AlipayCertPublicKey string `yaml:"alipay_cert_public_key"`
	EncryptKey          string `yaml:"encrypt_key"`
	ReturnURL           string `yaml:"return_url"`
	NotifyURL           string `yaml:"notify_url"`
	Server              string `yaml:"server"`
	Release             bool   `yaml:"release"`
}

type Paypal struct {
	ClientID  string `yaml:"client_id"`
	Secret    string `yaml:"secret"`
	ReturnURL string `yaml:"return_url"`
	CancelURL string `yaml:"cancel_url"`
	NotifyURL string `yaml:"notify_url"`
	Server    string `yaml:"server"`
}
