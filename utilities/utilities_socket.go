package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	// ConnWS is array of websocket created
	ConnWS []*websocket.Conn
	// MessageType is ...
	MessageType int
	// upgrader is ...
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	muTx = &sync.Mutex{}
)

// SocketMessage is Response type of socket message
type SocketMessage struct {
	Channel string
	Message string
}

// reader is to read socket values
func reader(connWS *websocket.Conn) {
	ConnWS = append(ConnWS, connWS)
	for {
		// Read message from browser
		msgType, msg, err := connWS.ReadMessage()
		MessageType = msgType
		if err != nil {
			for i, connection := range ConnWS {
				if connection == connWS {
					ConnWS = remove(ConnWS, i)
					break
				}
			}
			return
		}
		fmt.Println(len(ConnWS))
		if err = connWS.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

// remove is to remove socket connect
func remove(s []*websocket.Conn, i int) []*websocket.Conn {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// Initialize is to initialize socket connection
func Initialize(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	if err != nil {
		fmt.Println("Error with socket")
	}
	reader(ws)

}

// Notify is to write on socket message
func Notify(message string) {
	msg := []byte(message)
	for _, conn := range ConnWS {
		muTx.Lock()
		err := conn.WriteMessage(1, msg)
		muTx.Unlock()
		if err != nil {
			fmt.Println("Error")
		}
	}
}

//SendNotification is to send notification
func SendNotification(code string, message string) {
	bytes, _ := json.Marshal(SocketMessage{
		code,
		message,
	})

	Notify(string(bytes))
}
