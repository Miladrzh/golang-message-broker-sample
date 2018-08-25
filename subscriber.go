package main

import (
	"fmt"
	"time"
)

// Subscriber declaration and methods
type Subscriber struct {
	queue MessageQueue
}

func (subscriber Subscriber) listen() {
	for {
		select {
		case message := <-subscriber.queue.receive():
			fmt.Printf("\n\nreceived %v from %v at time %v", message.content, message.sender, time.Now())
		default:
			time.Sleep(1000000000)
		}
	}
}

// end of Subscriber
