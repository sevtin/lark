package xs3

type PresignedUrlInput struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
}

type PresignedUrlOutput struct {
	PutUrl      string `json:"put_url"`
	Url         string `json:"url"`
	Key         string `json:"key"`
	ContentType string `json:"content_type"`
	Acl         string `json:"acl"`
	Filename    string `json:"filename"`
}
