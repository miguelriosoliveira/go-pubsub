package pubsub

type Publisher interface {
	AddSubscriber(subscriber Subscriber)
	Publish(message Message)
}

type publisher struct {
	subscribers []Subscriber
	messages    []Message
}

func NewPublisher() Publisher {
	return &publisher{
		subscribers: []Subscriber{},
		messages:    []Message{},
	}
}

func (p *publisher) AddSubscriber(subscriber Subscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p *publisher) Publish(message Message) {
	p.messages = append(p.messages, message)
	for _, sub := range p.subscribers {
		sub.Consume(message)
	}
}
