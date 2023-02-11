package obj

type KafkaMessageHandler func(msg []byte, msgKey string) (err error)
