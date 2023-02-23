package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	SERVER := "localhost:8080"
	PATH := "/"

	URL := url.URL{Scheme: "ws", Host: SERVER, Path: PATH}

	log.Println("Connecting to: ", URL.String())

	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("ReadMessage() error:", err)
			return
		}
		log.Printf("Received: %s", message)
	}
}
