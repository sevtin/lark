package service

type ChatGptService interface {
	Run()
}

type chatGptService struct {
}

func NewChatGptService() ChatGptService {
	return &chatGptService{}
}

func (s *chatGptService) Run() {

}
