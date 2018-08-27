package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

type MessageBroker struct {
	queues map[string]MessageQueue
}

type Message struct {
	Content string `json:"message"`
	Time    string `json:"time"`
}

func (broker MessageBroker) listen() {
	http.HandleFunc("/publish/", broker.handleRequest)
	http.HandleFunc("/subscribe/", broker.handleRequest)
	panic(http.ListenAndServe(":8000", nil))
}

func (broker MessageBroker) handleRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open web socket connection", http.StatusBadRequest)
	}

	queueName := strings.Replace(r.URL.Path, "/publish/", "", 1)
	queueName = strings.Replace(queueName, "/subscribe/", "", 1)
	mQueue := broker.getMessageQueue(queueName)

	if r.URL.Path[:9] == "/publish/" {
		go mQueue.listenToPublisher(conn, queueName)
	} else if r.URL.Path[:11] == "/subscribe/" {
		go mQueue.listenToSubscriber(conn, queueName)
	}
}

func (broker MessageBroker) getMessageQueue(queueName string) (MessageQueue) {
	_, ok := broker.queues[queueName]
	if ok == false {
		mQueue := MessageQueue{make(map[int]chan Message)}
		broker.queues[queueName] = mQueue
	}
	return broker.queues[queueName]
}

func main() {
	broker := MessageBroker{make(map[string]MessageQueue)}
	broker.listen()
}
