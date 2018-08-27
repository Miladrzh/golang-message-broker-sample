package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

type Message struct {
	Content string `json:"message"`
	Time    string `json:"time"`
}

func GetWebSocketFromURL(scheme string, hostAddress string, path string) (*websocket.Conn) {
	wsURL := url.URL{Scheme: scheme, Host: hostAddress, Path: path}
	ws, _, err := websocket.DefaultDialer.Dial(wsURL.String(), nil)

	log.Printf("connecting to %s", wsURL.String())
	if err != nil {
		log.Fatal("dial:", err)
	}
	return ws
}
