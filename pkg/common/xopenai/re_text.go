package xopenai

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatMessageReq struct {
	Model            string     `json:"model"`
	Messages         []*Message `json:"messages"`
	MaxTokens        int        `json:"max_tokens"`
	Temperature      float32    `json:"temperature"`
	TopP             int        `json:"top_p"`
	FrequencyPenalty int        `json:"frequency_penalty"`
	PresencePenalty  int        `json:"presence_penalty"`
}

type ChatMessageResp struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Choices []*Choice `json:"choices"`
	Usage   *Usage    `json:"usage"`
}

type Choice struct {
	Message      *Message `json:"message"`
	Index        int      `json:"index"`
	FinishReason string   `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
