package main

import "github.com/miguelriosoliveira/go-pubsub/pubsub"

func main() {
	pub := pubsub.NewPublisher()

	sub1 := pubsub.NewSubscriber("sub1")
	sub2 := pubsub.NewSubscriber("sub2")

	pub.AddSubscriber(sub1)
	pub.AddSubscriber(sub2)

	pub.Publish(pubsub.NewMessage("hello!"))
	pub.Publish(pubsub.NewMessage("world!"))
}
