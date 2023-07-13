package xopenai

type ImageVariationReq struct {
	Image          string `json:"image"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
}

type ImageVariationResp struct {
	Created int64 `json:"created"`
	Data    []struct {
		Base64Json string `json:"b64_json"`
	} `json:"data"`
}

// 图片生成
type ImageGenerateReq struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
}

type ImageGenerateResp struct {
	Created int64 `json:"created"`
	Data    []struct {
		Base64Json string `json:"b64_json"`
	} `json:"data"`
}
