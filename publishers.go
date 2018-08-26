package main

import (
	"fmt"
	"time"
)

// implementation of first publisher declared in phase#2 description
func publishType1(mQueue MessageQueue) {
	for {
		for i := 0; i < 10; i++ {
			mQueue.send(Message{"publisher1", i})
			fmt.Printf("\n\npublisher 1 at time %v the message is:\n%v", time.Now(), i)
		}
		time.Sleep(1000000000)
	}
}

// implementation of second publisher declared in phase#2 description
func publishType2(mQueue MessageQueue) {
	for {
		mQueue.send(Message{"publisher2", "hello"})
		fmt.Printf("\n\npublisher 2 at time %v the message is:\nhello", time.Now())
		time.Sleep(500000000)
	}
}

// implementation of third publisher declared in phase#2 description
func publishType3(mQueue MessageQueue) {
	for {
		for i := 20; i >= 0; i-- {
			mQueue.send(Message{"publisher3", i})
			fmt.Printf("\n\npublisher 1 at time %v the message is:\n%v", time.Now(), i)
		}
		time.Sleep(700000000)
	}
}
