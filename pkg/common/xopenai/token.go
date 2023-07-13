package xopenai

type ChatGptToken struct {
	tokens   []string
	length   int
	curIndex int
}

func NewChatGptToken(tokens []string) *ChatGptToken {
	return &ChatGptToken{
		tokens: tokens,
		length: len(tokens),
	}
}

func (c *ChatGptToken) GetToken() (token string) {
	c.curIndex = (c.curIndex + 1) % c.length
	token = c.tokens[c.curIndex]
	return
}
