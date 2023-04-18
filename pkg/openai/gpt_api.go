package openai

import (
	"fmt"
)

const (
	GPT_MAX_TOKENS  = 2000
	GPT_TEMPERATURE = 0.7
	GPT_ENGINE      = "gpt-3.5-turbo"
)

const (
	GPT_API_CHAT_COMPLETIONS     = iota // 聊天消息
	GPT_API_AUDIO_TRANSCRIPTIONS        // 音频转文本
	GPT_API_IMAGES_VARIATIONS
	GPT_API_IMAGES_GENERATIONS // 生成一张图片
)

func (gpt *chatGpt) FullUrl(apiType int) (url string) {
	switch apiType {
	case GPT_API_CHAT_COMPLETIONS:
		url = fmt.Sprintf("%s/v1/chat/completions", gpt.ApiUrl)
	case GPT_API_AUDIO_TRANSCRIPTIONS:
		url = fmt.Sprintf("%s/v1/audio/transcriptions", gpt.ApiUrl)
	case GPT_API_IMAGES_VARIATIONS:
		url = fmt.Sprintf("%s/v1/images/variations", gpt.ApiUrl)
	case GPT_API_IMAGES_GENERATIONS:
		url = fmt.Sprintf("%s/v1/images/generations", gpt.ApiUrl)
	}
	return
}

// 文本消息
func (gpt *chatGpt) SendTextMessage(msgs []*Message) (resp *ChatMessageResp, err error) {
	var (
		req = &ChatMessageReq{
			Model:            GPT_ENGINE,
			Messages:         msgs,
			MaxTokens:        GPT_MAX_TOKENS,
			Temperature:      GPT_TEMPERATURE,
			TopP:             1,
			FrequencyPenalty: 0,
			PresencePenalty:  0,
		}
	)
	resp = new(ChatMessageResp)
	err = gpt.sendRequest(gpt.FullUrl(GPT_API_CHAT_COMPLETIONS), "POST", GPT_BODY_TYPE_JSON, req, resp)
	return
}

// 声音转文字
func (gpt *chatGpt) AudioToText(file string) (resp *AudioToTextResp, err error) {
	var (
		req = &AudioToTextReq{
			File:           file,
			Model:          "whisper-1",
			ResponseFormat: "text",
		}
	)
	resp = &AudioToTextResp{}
	err = gpt.sendRequest(gpt.FullUrl(GPT_API_AUDIO_TRANSCRIPTIONS), "POST", GPT_BODY_TYPE_VOICE, req, resp)
	return
}

// 创建一个图片
func (gpt *chatGpt) ImageGenerate(prompt string, size string) (resp *ImageGenerateResp, err error) {
	var (
		req = &ImageGenerateReq{
			Prompt:         prompt,
			N:              1,
			Size:           size,
			ResponseFormat: "b64_json",
		}
	)
	resp = new(ImageGenerateResp)
	err = gpt.sendRequest(gpt.FullUrl(GPT_API_IMAGES_GENERATIONS), "POST", GPT_BODY_TYPE_JSON, req, resp)
	return
}

func (gpt *chatGpt) ImageVariation(image string, size string) (resp *ImageVariationResp, err error) {
	var (
		req = &ImageVariationReq{
			Image:          image,
			N:              1,
			Size:           size,
			ResponseFormat: "b64_json",
		}
	)
	resp = &ImageVariationResp{}
	err = gpt.sendRequest(gpt.FullUrl(GPT_API_IMAGES_VARIATIONS), "POST", GPT_BODY_TYPE_PICTURE, req, resp)
	return
}
