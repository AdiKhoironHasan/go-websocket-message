package wsgo

type SocketReqDTO struct {
	From    string  `json:"from"`
	ForUser string  `json:"for_user"`
	Message Message `json:"message"`
}

type Message struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// type WebSocketConnection struct {
// 	Conn *websocket.Conn
// 	Username string
// }

// type SocketResponse struct {
// 	From    string
// 	ForUser string // spesifict user
// 	Type    string
// 	Message string
// }

// type SimpleWebsocket interface{
// 	broadcastMessage(currentConn *WebSocketConnection, kind, message string)
// }

// type simpleWebsocket struct{
// 	Connection *websocket.Conn
// 	Username string
// }

// func NewSimpleWebsocket(Connection *websocket.Conn,Username string) SimpleWebsocket {
// 	return &simpleWebsocket{
// 		Connection: Connection,
// 		Username: Username,
// 	}
// }

// func BroadcastMessage(connections *[]WebSocketConnection, currentConn *WebSocketConnection, kind, message string) {
// 	for _, eachConn := range *connections {
// 		if eachConn == *currentConn {
// 			continue
// 		}

// 		eachConn.Conn.WriteJSON(SocketResponse{
// 			From:    currentConn.Username,
// 			Type:    kind,
// 			Message: message,
// 		})
// 	}
// }
