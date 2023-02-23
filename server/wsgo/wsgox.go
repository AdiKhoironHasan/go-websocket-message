package wsgo

// import (
// 	"fmt"
// 	"log"
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

// // func MakeConn(w http.ResponseWriter, r *http.Request, responseHeader http.Header, readBufSize, writeBufSize int) (*websocket.Conn, error) {
// // 	currentGorillaConn, err := websocket.Upgrade(w, r, responseHeader, readBufSize, writeBufSize)
// // 	if err != nil {
// // 			return nil, errors.New("could not open websocket connection")
// // 			// http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
// // 	}

// // 	return currentGorillaConn, nil
// // }

// // func GetCurrentConn(Conn *websocket.Conn, Id string) WebSocketConnection {
// // 	return WebSocketConnection{Conn: Conn, Username: Id}
// // }

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
