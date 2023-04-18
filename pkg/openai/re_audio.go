package openai

type AudioToTextReq struct {
	File           string `json:"file"`
	Model          string `json:"model"`
	ResponseFormat string `json:"response_format"`
}

type AudioToTextResp struct {
	Text string `json:"text"`
}
