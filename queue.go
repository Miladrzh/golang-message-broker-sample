package main

import "math/rand"

type Message struct {
	sender  string
	content interface{}
}

// MessageQueue declaration and methods
type MessageQueue struct {
	channels map[int]chan Message
}

func (mQueue MessageQueue) registerSubscriber() (int) {
	key := rand.Intn(1000000)
	mQueue.channels[key] = make(chan Message, 10)
	return key
}

func (mQueue MessageQueue) send(message Message) {
	for key := range mQueue.channels {
		mQueue.channels[key] <- message
	}
}

func (mQueue MessageQueue) receive(key int) (chan Message) {
	return mQueue.channels[key]
}

// end of MessageQueue
