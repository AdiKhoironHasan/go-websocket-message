package main

import (
	"dms/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/adikhoironhasan/gowes"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type SocketPayload struct {
// 	Message string `json:"message"`
// }

type SocketResponse struct {
	From    string `json:"from"`
	ForUser string `json:"for_user"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

var DB *gorm.DB

func main() {
	// database connection
	dsn := "root@tcp(localhost:3306)/ws_notifications?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error:", err)
		return
	}
	DB = db

	SERVER := "localhost:8080"
	PATH := "/"

	fmt.Println("Connecting to ws server:", SERVER, "at", PATH)

	URL := url.URL{
		Scheme:   "ws",
		Host:     SERVER,
		Path:     PATH,
		RawQuery: "user=dms",
	}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("ReadMessage() error:", err)
				return
			}
			// log.Printf("Received: %s", message)

			messageData := gowes.SocketPayload{}
			// unmarshal message
			err = json.Unmarshal(message, &messageData)
			if err != nil {
				panic(err)
			}

			// for update notification as read in database
			if messageData.MessageType == gowes.TypeReadNotification {
				notificationID := messageData.MessageData.(map[string]interface{})["notification_id"].(float64)
				// fmt.Println(notificationID)
				DB.Model(&entity.NotificationEntity{}).Select("is_read").Where("id = ?", notificationID).Updates(map[string]interface{}{"is_read": true})
			}
			// for read notification from database
			if messageData.MessageType == gowes.TypeGetNotification {
				user, err := strconv.Atoi(strings.SplitAfter(messageData.To[0].ClientID, "user-")[1])
				if err != nil {
					panic(err)
				}
				// fmt.Println(user)

				dataNotification := []entity.NotificationEntity{}
				rows := DB.Model(&entity.NotificationEntity{}).Where("user_id = ?", user).Scan(&dataNotification)
				if rows.Error != nil {
					panic(rows.Error)
				}

				if len(dataNotification) > 0 {
					// messageNotification := []gowes.Message{}
					// for _, each := range dataNotification {
					// 	// fmt.Println(each.Detail)
					// 	messageNotification = append(messageNotification, gowes.Message{
					// 		Title:  each.Title,
					// 		Detail: each.Detail,
					// 	})
					// }

					// send notification via websocket
					c.WriteJSON(gowes.SocketPayload{
						To: []gowes.Client{
							{ClientID: "user-" + strconv.Itoa(user)},
						},
						MessageType: gowes.TypeSendNotification,
						MessageData: dataNotification,
					})
				}
			}

		}
	}()

	mux := mux.NewRouter()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to dms")
	}).Methods("GET")

	mux.HandleFunc("/asign", func(w http.ResponseWriter, r *http.Request) {
		// get data from request
		data := gowes.SocketPayload{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// prepare data for notification table
		notificationData := []entity.NotificationParam{}
		for _, val := range data.To {
			username, err := strconv.Atoi(strings.SplitAfter(val.ClientID, "user-")[1])
			if err != nil {
				panic(err)
			}
			dataMessage := data.MessageData.(map[string]interface{})
			notificationData = append(notificationData, entity.NotificationParam{
				UserID: username,
				Title:  dataMessage["title"].(string),
				Detail: dataMessage["detail"].(string),
			})
		}

		// insert into notification table
		DB.Create(&notificationData)
		log.Println("success insert into notification table")

		data.MessageType = gowes.TypeSendNotification
		// send notification via websocket
		c.WriteJSON(data)
	}).Methods("POST")

	fmt.Println("Server starting at :8082")
	http.ListenAndServe(":8082", mux)
}
