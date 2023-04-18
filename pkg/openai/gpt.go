package openai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

type PlatForm string

const (
	GPT_BODY_TYPE_JSON = iota
	GPT_BODY_TYPE_VOICE
	GPT_BODY_TYPE_PICTURE
)

const (
	GPT_REQUEST_MAX_RETRIES = 3
)

type ChatGpt interface {
	SendTextMessage(msgs []*Message) (resp *ChatMessageResp, err error)
	AudioToText(audio string) (resp *AudioToTextResp, err error)
	ImageGenerate(prompt string, size string) (resp *ImageGenerateResp, err error)
	ImageVariation(images string, size string) (resp *ImageVariationResp, err error)
}

type chatGpt struct {
	token  *ChatGptToken
	ApiUrl string
}

func NewChatGpt(tokens []string, apiUrl string) ChatGpt {
	gpt := &chatGpt{}
	gpt.token = NewChatGptToken(tokens)
	gpt.ApiUrl = apiUrl
	return gpt
}

func (gpt *chatGpt) sendRequest(url, method string, bodyType int, requestBody interface{}, responseBody interface{}) (err error) {
	var (
		reqBody []byte
		writer  *multipart.Writer
		request *http.Request
		data    []byte
		client  = &http.Client{Timeout: 120 * time.Second}
	)
	switch bodyType {
	case GPT_BODY_TYPE_JSON:
		if reqBody, err = json.Marshal(requestBody); err != nil {
			return
		}
	case GPT_BODY_TYPE_VOICE:
		buffer := &bytes.Buffer{}
		writer = multipart.NewWriter(buffer)
		if err = audioMultipart(requestBody.(*AudioToTextReq), writer); err != nil {
			return
		}
		if err = writer.Close(); err != nil {
			return
		}
		reqBody = buffer.Bytes()
	case GPT_BODY_TYPE_PICTURE:
		buffer := &bytes.Buffer{}
		writer = multipart.NewWriter(buffer)
		if err = pictureMultipart(requestBody.(*ImageVariationReq), writer); err != nil {
			return
		}
		if err = writer.Close(); err != nil {
			return
		}
		reqBody = buffer.Bytes()
	default:
		return
	}
	if request, err = http.NewRequest(method, url, bytes.NewReader(reqBody)); err != nil {
		return
	}
	if bodyType == GPT_BODY_TYPE_VOICE || bodyType == GPT_BODY_TYPE_PICTURE {
		request.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("Authorization", "Bearer "+gpt.token.GetToken())

	var (
		response *http.Response
		retry    int
	)
	for retry = 0; retry <= GPT_REQUEST_MAX_RETRIES; retry++ {
		response, err = client.Do(request)
		if err != nil || response.StatusCode < 200 || response.StatusCode >= 300 {
			if retry == GPT_REQUEST_MAX_RETRIES {
				break
			}
			time.Sleep(time.Duration(retry+1) * time.Second)
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
	}
	if response == nil || response.StatusCode < 200 || response.StatusCode >= 300 {
		return
	}
	if data, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	if err = json.Unmarshal(data, responseBody); err != nil {
		return
	}
	return
}
