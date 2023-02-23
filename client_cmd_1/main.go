package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From    string
	ForUser string
	Type    string
	Message string
}

func main() {

	SERVER := "localhost:8080"
	PATH := "/"

	fmt.Println("Connecting to:", SERVER, "at", PATH)

	URL := url.URL{Scheme: "ws", Host: SERVER, Path: PATH}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	defer c.Close()

	num := 1
	for {
		// infinity dummy data
		messageData := fmt.Sprintf("Hello World %d", num)
		responseData := SocketResponse{
			From: "client",
			Type: "message",
			// ForUser: "User 4", // spesific user
			Message: messageData,
		}
		c.WriteJSON(responseData)
		if err != nil {
			log.Println("Write error:", err)
			return
		}
		fmt.Println("Send:", messageData)
		num++
		time.Sleep(time.Second * 3)
	}
}
