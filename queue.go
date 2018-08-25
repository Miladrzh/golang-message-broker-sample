package main

type Message struct {
	sender  string
	content interface{}
}

// MessageQueue declaration and methods
type MessageQueue struct {
	channel chan Message
}

func (mQueue MessageQueue) send(message Message) {
	mQueue.channel <- message
}

func (mQueue MessageQueue) receive() (chan Message) {
	return mQueue.channel
}

// end of MessageQueue
