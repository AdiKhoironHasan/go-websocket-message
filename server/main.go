package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/adikhoironhasan/gowes"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

const MESSAGE_NEW_USER = "New User"
const MESSAGE_CHAT = "Chat"
const MESSAGE_LEAVE = "Leave"

var connections = make([]*WebSocketConnection, 0)

type SocketPayload struct {
	Message string
	ForUser string // spesifict user
}

type SocketResponse struct {
	From    string
	ForUser []string // spesifict user
	Type    string
	Message string
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}

		username := r.URL.Query().Get("user")

		currentConn := WebSocketConnection{
			Conn:     currentGorillaConn,
			Username: username,
		}

		if currentConn.Conn != nil {
			connections = append(connections, &currentConn)

			if currentConn.Username != "dms" {
				// send message to get notification from dms
				sendPrivateMessage(&currentConn, gowes.TypeGetNotification, currentConn.Username, gowes.TypeGetNotification)
			}
		}

		go handleIO(&currentConn, connections)
	})

	mux.HandleFunc("/get-user", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(connections)
		w.Write(data)
	})

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", mux)
}

func handleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	// broadcastMessage(currentConn, MESSAGE_NEW_USER, "")
	for {
		payload := gowes.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		// fmt.Println("lanjut")

		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				// send meesage user left
				// broadcastMessage(currentConn, MESSAGE_LEAVE, "")
				ejectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		// for private user
		// if payload.ForUser != "" {
		// add to package
		// if message type is send notification to browser
		if payload.MessageType == gowes.TypeSendNotification {
			sendNotification(currentConn, payload.MessageData, payload.To)
		}
		if payload.MessageType == gowes.TypeReadNotification {
			readNotification("dms", payload.MessageType, payload.MessageData)
		}
		// privateMessage(currentConn, MESSAGE_CHAT, payload.Message.Title, payload.ForUser)
		// } else {
		// 	broadcastMessage(currentConn, MESSAGE_CHAT, payload.Message)
		// }

	}
}

func ejectConnection(currentConn *WebSocketConnection) {
	filtered := gubrak.From(connections).Reject(func(each *WebSocketConnection) bool {
		return each == currentConn
	}).Result()
	connections = filtered.([]*WebSocketConnection)
}

// func broadcastMessage(currentConn *WebSocketConnection, kind, message string) {
// 	for _, eachConn := range connections {
// 		if eachConn == currentConn {
// 			continue
// 		}

// 		eachConn.WriteJSON(SocketResponse{
// 			From:    currentConn.Username,
// 			Type:    kind,
// 			Message: message,
// 		})
// 	}
// }

// func privateMessage(currentConn *WebSocketConnection, kind, message string, user string) {
// 	for _, eachConn := range connections {
// 		// fmt.Println("ini loop ke - ", k)
// 		fmt.Println("ini usernya: ", eachConn.Username)
// 		if eachConn == currentConn {
// 			continue
// 		}

// 		if eachConn.Username == user {
// 			eachConn.WriteJSON(SocketResponse{
// 				From:    currentConn.Username,
// 				ForUser: user,
// 				Type:    kind,
// 				Message: message,
// 			})
// 		}
// 	}
// }

func readNotification(user, typeMessage string, message any) {
	for _, val := range connections {
		if val.Username == user {
			val.Conn.WriteJSON(gowes.SocketPayload{
				MessageData: message,
				MessageType: typeMessage,
			})
		}
	}

}

func sendPrivateMessage(currentConn *WebSocketConnection, message, user, typeMessage string) {

	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		// send to dms
		if eachConn.Username == "dms" {
			eachConn.Conn.WriteJSON(gowes.SocketPayload{
				MessageData: message,
				MessageType: typeMessage,
				To: []gowes.Client{
					{
						ClientID: user,
					},
				},
			})
		}
	}
}

func sendNotification(currentConn *WebSocketConnection, message any, users []gowes.Client) {
	// fmt.Println(connections)
	for _, user := range users {
		for _, eachConn := range connections {
			if eachConn == currentConn {
				continue
			}

			// fmt.Println("ini loop ke - ", k)
			if eachConn.Username == user.ClientID {
				eachConn.Conn.WriteJSON(gowes.SocketPayload{
					MessageData: message,
					MessageType: gowes.TypeSendNotification,
				})
				// fmt.Println("kirim ke : ", eachConn.Username)
			}
			// else {
			// fmt.Println("skip user : ", eachConn.Username)
			// }
		}
	}

	// if currentConn.Username == "dms" {
	// 	fmt.Println("send: ", gowes.SocketPayload{
	// 		MessageData: message,
	// 		To:          users,
	// 	})
	// 	// send data (ws server > dms) to save data notification to database
	// 	currentConn.Conn.WriteJSON(gowes.SocketPayload{
	// 		MessageData: message,
	// 		To:          users,
	// 	})
	// }
}
