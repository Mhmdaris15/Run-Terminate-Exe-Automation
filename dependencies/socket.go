package dependencies

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type hub struct {
	clients               map[*websocket.Conn]bool
	clientRegisterChannel chan *websocket.Conn
	clientRemovalChannel  chan *websocket.Conn
	broadcastMessage      chan Message
}

func (h *hub) SocketRun() {
	for {
		select {
		case conn := <-h.clientRegisterChannel:
			h.clients[conn] = true
		case conn := <-h.clientRemovalChannel:
			delete(h.clients, conn)
		case message := <-h.broadcastMessage:
			for conn := range h.clients {
				err := conn.WriteJSON(message)
				if err != nil {
					log.Println("Error sending message:", err)
				}
			}
		}
	}
}

func SocketInitiate() *hub {
	h := &hub{
		clients:               make(map[*websocket.Conn]bool),
		clientRegisterChannel: make(chan *websocket.Conn),
		clientRemovalChannel:  make(chan *websocket.Conn),
		broadcastMessage:      make(chan Message),
	}

	return h
}

func SocketAllowUpgrade(h *hub) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		h.clientRegisterChannel <- conn
	}

}

func BidScore(h *hub) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		defer func() {
			h.clientRemovalChannel <- conn
			conn.Close()
		}()

		h.clientRegisterChannel <- conn

		for {
			messageType, payload, err := conn.ReadMessage()
			if err != nil {
				// if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				// 	log.Println("Error reading message:", err)
				// } else {
				// 	log.Println("Error reading message:", err)
				// }
				return
			}

			if messageType == websocket.TextMessage {
				var message Message
				err := json.Unmarshal(payload, &message)
				if err != nil {
					log.Print("Error decoding JSON message: ", err)
					return
				}

				// Check if type is bid
				if message.Type == "bid" {
					h.broadcastMessage <- Message{
						Type:    message.Type,
						Payload: message.Payload,
					}

				} else if message.Type == "delegate" {
					h.broadcastMessage <- Message{
						Type:    message.Type,
						Payload: message.Payload,
					}
				} else if message.Type == "run" {

					ws := message.Payload.(map[string]interface{})["ws"].(string)
					log.Printf("Running %s", ws)

					wsPath, err := ReadWs(ws)
					if err != nil {
						log.Println("Error reading ws:", err)
						return
					}
					log.Printf("Path: %s", wsPath)

					// Check if exe is already running or not
					_, ok := RunningExes[wsPath]
					if ok {
						log.Printf("Exe %s is already running", wsPath)
						h.broadcastMessage <- Message{
							Type:    "already_running",
							Payload: message.Payload,
						}
						return
					}

					err = RunningExe(wsPath, ws)
					if err != nil {
						log.Println("Error running exe:", err)
						return
					}

					h.broadcastMessage <- Message{
						Type:    message.Type,
						Payload: message.Payload,
					}
				} else if message.Type == "terminate" {

					ws := message.Payload.(map[string]interface{})["ws"].(string)
					log.Printf("Terminating %s", ws)

					wsPath, err := ReadWs(ws)
					if err != nil {
						log.Println("Error reading ws:", err)
						return
					}
					log.Printf("Path: %s", wsPath)

					err = TerminateExe(wsPath)
					if err != nil {
						log.Println("Error terminating exe:", err)
						return
					}

					h.broadcastMessage <- Message{
						Type:    message.Type,
						Payload: message.Payload,
					}
				}
			}
		}
	}
}
