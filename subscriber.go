package main

import (
	"MessageQueue/client"
	"fmt"
)

func main() {
	var queueName string
	fmt.Println("Enter queue name:")
	fmt.Scan(&queueName)

	ws := client.GetWebSocketFromURL("ws", "localhost:8000", "/subscribe/"+queueName)

	for {
		message := client.Message{}
		err := ws.ReadJSON(&message)
		if err == nil {
			fmt.Printf("received    %v\n", message)
		}
	}
}
