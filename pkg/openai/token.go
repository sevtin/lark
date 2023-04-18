package openai

import "sync"

type ChatGptToken struct {
	tokens   []string
	length   int
	curIndex int
	mutex    sync.RWMutex
}

func NewChatGptToken(tokens []string) *ChatGptToken {
	return &ChatGptToken{
		tokens: tokens,
		length: len(tokens),
	}
}

func (c *ChatGptToken) GetToken() (token string) {
	c.mutex.RLock()
	c.curIndex = (c.curIndex + 1) % c.length
	token = c.tokens[c.curIndex]
	c.mutex.RUnlock()
	return
}
