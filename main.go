package main

import (
	"time"
)

func main() {
	mQueue := MessageQueue{make(chan Message)}

	subscriber := Subscriber{mQueue}
	go subscriber.listen()

	go publishType1(mQueue)
	go publishType2(mQueue)
	go publishType3(mQueue)

	for {
		time.Sleep(10000000000)
	}
}
