package main

import (
	"MessageQueue/client"
	"fmt"
)

func main() {
	var queueName string
	fmt.Println("publisher connecting to sample queue")
	queueName = "sample_queue"

	ws := client.GetWebSocketFromURL("ws", "broker:8000", "/publish/"+queueName)

	for {
		var content string
		fmt.Println("publisher send sample message")
		content = "sample message"
		message := client.Message{}
		message.Content = content
		ws.WriteJSON(message)
	}
}
