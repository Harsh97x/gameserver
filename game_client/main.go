package main

import (
	"log"

	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws//:localhost:40000/ws"

func main() {
	dialer := websocket.Dialer{ReadBufferSize: 1024, WriteBufferSize: 1024}
	conn, _, err := dialer.Dial(wsServerEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
	}
}
