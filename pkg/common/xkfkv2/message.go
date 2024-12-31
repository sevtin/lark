package xkafka

type Message struct {
	Id    int64  `json:"id"`
	Topic string `json:"topic"`
}
