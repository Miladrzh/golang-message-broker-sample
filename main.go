package main

import (
	"fmt"
	"time"
)

type Message struct {
	content interface{}
}

// MessageQueue declaration and methods
type MessageQueue struct {
	channel chan Message
}

func (mQueue MessageQueue) send(message Message) {
	mQueue.channel <- message
}

func (mQueue MessageQueue) receive() (r Message) {
	r = <-mQueue.channel
	return
}

// end of MessageQueue

// Publisher declaration and methods
type Publisher struct {
	queue MessageQueue
}

func (publisher Publisher) publish() {
	for {
		for i := 0; i < 10; i++ {
			publisher.queue.send(Message{i})
			time.Sleep(1000000000)
		}
		fmt.Print("\n")
	}
}

// end of Publisher

// Subscriber declaration and methods
type Subscriber struct {
	queue MessageQueue
}

func (subscriber Subscriber) listen() {
	for {
		message := subscriber.queue.receive()
		fmt.Printf("%v", message)
	}
}

// end of Subscriber

func main() {
	mQueue := MessageQueue{make(chan Message)}
	subscriber := Subscriber{mQueue}
	publisher := Publisher{mQueue}
	go publisher.publish()
	go subscriber.listen()
	for {
		time.Sleep(10000000000)
	}
}
