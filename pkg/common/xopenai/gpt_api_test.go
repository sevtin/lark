package xopenai

import (
	"fmt"
	"testing"
)

const (
	GPT_TOKEN = "sk-"
	GPT_API   = "https://api.openai.com"
)

func TestSendMessage(t *testing.T) {
	var (
		gpt = NewChatGpt([]string{GPT_TOKEN}, GPT_API)
		msg = &Message{
			Role:    "user",
			Content: "香港的全名",
		}
		resp *ChatMessageResp
		err  error
	)
	resp, err = gpt.SendTextMessage([]*Message{msg})
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.Created == 0 {

	}
}

func TestAudioToText(t *testing.T) {
	var (
		gpt  = NewChatGpt([]string{GPT_TOKEN}, GPT_API)
		resp *AudioToTextResp
		err  error
	)
	resp, err = gpt.AudioToText("./openai/test_file/test.wav")
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.Text != "" {

	}
}

func TestGenerateImage(t *testing.T) {
	var (
		gpt  = NewChatGpt([]string{GPT_TOKEN}, GPT_API)
		resp *ImageGenerateResp
		err  error
	)
	resp, err = gpt.ImageGenerate("苹果", "512x512")
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.Created != 0 {

	}
}

// 失败
func TestImagesVariation(t *testing.T) {
	var (
		gpt  = NewChatGpt([]string{GPT_TOKEN}, GPT_API)
		resp *ImageVariationResp
		err  error
	)
	resp, err = gpt.ImageVariation("./openai/test_file/img.png", "256x256")
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.Created != 0 {

	}
}

func TestGetToken(t *testing.T) {
	ct := NewChatGptToken([]string{"1", "2", "3", "4", "5"})
	for i := 0; i < 50; i++ {
		go fmt.Println(ct.GetToken())
	}
}
