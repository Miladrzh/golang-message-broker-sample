package main

import (
	"time"
)

func main() {
	mQueue := MessageQueue{make(map[int]chan Message)}

	key1 := mQueue.makeChannel()
	key2 := mQueue.makeChannel()
	subscriber1 := Subscriber{mQueue, key1}
	subscriber2 := Subscriber{mQueue, key2}

	go subscriber1.listen()
	go subscriber2.listen()

	go publishType1(mQueue, key1)
	go publishType2(mQueue, key1)
	go publishType3(mQueue, key2)

	for {
		time.Sleep(10000000000)
	}
}
