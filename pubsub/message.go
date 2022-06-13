package pubsub

type Message struct {
	body string
}

func NewMessage(text string) Message {
	return Message{body: text}
}
