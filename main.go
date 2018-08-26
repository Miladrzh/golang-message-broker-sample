package main

import (
	"time"
)

func main() {
	mQueue1 := MessageQueue{make(map[int]chan Message)}
	mQueue2 := MessageQueue{make(map[int]chan Message)}

	subscriber1 := Subscriber{mQueue1, "mQueue1", mQueue1.registerSubscriber()}
	subscriber2 := Subscriber{mQueue2, "mQueue2", mQueue2.registerSubscriber()}
	subscriber3 := Subscriber{mQueue2, "mQueue2", mQueue2.registerSubscriber()}

	go subscriber1.listen()
	go subscriber2.listen()
	go subscriber3.listen()

	go publishType1(mQueue1)
	go publishType2(mQueue2)
	go publishType3(mQueue2)

	for {
		time.Sleep(10000000000)
	}
}
