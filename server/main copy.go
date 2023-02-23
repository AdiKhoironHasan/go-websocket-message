package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"github.com/gorilla/websocket"
// 	gubrak "github.com/novalagung/gubrak/v2"
// )

// type M map[string]interface{}

// const MESSAGE_NEW_USER = "New User"
// const MESSAGE_CHAT = "Chat"
// const MESSAGE_LEAVE = "Leave"

// var connections = make([]*WebSocketConnection, 0)

// type SocketPayload struct {
// 	Message string
// 	ForUser string // spesifict user
// }

// type SocketResponse struct {
// 	From    string
// 	ForUser string // spesifict user
// 	Type    string
// 	Message string
// }

// type WebSocketConnection struct {
// 	*websocket.Conn
// 	Username string
// }

// func main() {
// 	numUser := 1
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
// 		if err != nil {
// 			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
// 		}

// 		// username := r.URL.Query().Get("username")
// 		currentConn := WebSocketConnection{Conn: currentGorillaConn, Username: fmt.Sprintf("User %d", numUser)}
// 		if currentConn.Conn != nil {
// 			numUser++
// 			connections = append(connections, &currentConn)
// 		}

// 		go handleIO(&currentConn, connections)
// 	})

// 	http.HandleFunc("/get-user", func(w http.ResponseWriter, r *http.Request) {
// 		data, _ := json.Marshal(connections)
// 		w.Write(data)
// 	})

// 	fmt.Println("Server starting at :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// func handleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			log.Println("ERROR", fmt.Sprintf("%v", r))
// 		}
// 	}()

// 	broadcastMessage(currentConn, MESSAGE_NEW_USER, "")
// 	for {
// 		payload := SocketPayload{}
// 		err := currentConn.ReadJSON(&payload)
// 		// fmt.Println("lanjut")
// 		if err != nil {
// 			if strings.Contains(err.Error(), "websocket: close") {
// 				broadcastMessage(currentConn, MESSAGE_LEAVE, "")
// 				ejectConnection(currentConn)
// 				return
// 			}

// 			log.Println("ERROR", err.Error())
// 			continue
// 		}

// 		// for private user
// 		if payload.ForUser != "" {
// 			privateMessage(currentConn, MESSAGE_CHAT, payload.Message, payload.ForUser)
// 		} else {
// 			broadcastMessage(currentConn, MESSAGE_CHAT, payload.Message)
// 		}

// 	}
// }

// func ejectConnection(currentConn *WebSocketConnection) {
// 	filtered := gubrak.From(connections).Reject(func(each *WebSocketConnection) bool {
// 		return each == currentConn
// 	}).Result()
// 	connections = filtered.([]*WebSocketConnection)
// }

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
