package main

import (
	"MessageQueue/client"
	"fmt"
)

func main() {
	var queueName string
	fmt.Println("subscriber connecting to sample queue")
	queueName = "sample_queue"

	ws := client.GetWebSocketFromURL("ws", "broker:8000", "/subscribe/"+queueName)

	for {
		message := client.Message{}
		err := ws.ReadJSON(&message)
		if err == nil {
			fmt.Printf("received    %v\n", message)
		}
	}
}
