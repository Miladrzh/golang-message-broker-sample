package main

import (
	"MessageQueue/client"
	"fmt"
)

func main() {
	var queueName string
	fmt.Println("Enter queue name:")
	fmt.Scan(&queueName)

	ws := client.GetWebSocketFromURL("ws", "broker:8000", "/publish/"+queueName)

	for {
		var content string
		fmt.Println("Enter message (type ! for exit)")
		fmt.Scan(&content)
		if content == "!" {
			return
		}
		message := client.Message{}
		message.Content = content
		ws.WriteJSON(message)
	}
}
