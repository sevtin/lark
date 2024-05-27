package xs3

type PresignedUrlInput struct {
	Key         string `json:"key"`
	Directory   string `json:"directory"`
	ContentType string `json:"content_type"`
}

type PresignedUrlOutput struct {
	PutUrl      string `json:"put_url"`
	Url         string `json:"url"`
	Key         string `json:"key"`
	ContentType string `json:"content_type"`
	Acl         string `json:"acl"`
}
