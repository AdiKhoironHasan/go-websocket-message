package wsgo

type SocketPayload struct {
	From    string    `json:"from"`
	ForUser []ForUser `json:"for_user"`
	Type    string    `json:"type"`
	// Message Message   `json:"message"`
	DataMessage interface{} `json:"data_message"`
}

type ForUser struct {
	Username string `json:"username"`
}

type Message struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type NotificationPayload struct {
	ToUser string `json:"to_user"`
	NotificationData []Message `json:"notification_data"`
}

const (
	// TypeMessage is a message type
	TypeGetNotification  = "get_notification"
	TypeSendNotification = "send_notification"
)

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
