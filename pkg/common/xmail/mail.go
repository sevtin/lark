package xmail

import (
	"gopkg.in/gomail.v2"
	"lark/pkg/conf"
)

func NewDialer(cfg *conf.Email) *gomail.Dialer {
	return gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
}

func NewMessage(from string, to string, subject string, content string) *gomail.Message {
	var (
		m = gomail.NewMessage()
	)
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	return m
}
