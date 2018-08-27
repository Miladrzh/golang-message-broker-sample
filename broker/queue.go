package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

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
		fmt.Println(key)
		mQueue.channels[key] <- message
	}
}

func (mQueue MessageQueue) receive(key int) (chan Message) {
	return mQueue.channels[key]
}

func (mQueue MessageQueue) listenToPublisher(conn *websocket.Conn, queueName string) {
	for {
		message := Message{}
		err := conn.ReadJSON(&message)
		if err != nil {
			return
		}
		go mQueue.send(message)
	}
}

func (mQueue MessageQueue) listenToSubscriber(conn *websocket.Conn, queueName string) {
	subscriberChannelId := mQueue.registerSubscriber()
	for {
		message := <-mQueue.receive(subscriberChannelId)
		message.Time = time.Now().String()
		fmt.Println(message)
		conn.WriteJSON(message)
	}
}

// end of MessageQueue
