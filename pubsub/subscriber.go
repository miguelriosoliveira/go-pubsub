package pubsub

import "fmt"

type Subscriber interface {
	Consume(message Message)
}

type subscriber struct {
	id string
}

func NewSubscriber(ID string) Subscriber {
	return &subscriber{
		id: "sub1",
	}
}

func (s *subscriber) Consume(message Message) {
	fmt.Println(s.id, "reading message:", message.body)
}
