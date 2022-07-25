package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var counter int = 0
var clients = make(map[*websocket.Conn]string)
var usernames = make(map[string]*websocket.Conn)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Target  string `json:"target"`
}

// {Part 1 }
func HandleClients(w http.ResponseWriter, r *http.Request) {
	go broadcastMessagesToClients()

	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket :: ", err)
	}
	defer websocket.Close()
	clients[websocket] = ""

	for {
		var message Message

		err := websocket.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred while reading message : %v", err)
			_, found := usernames[clients[websocket]]
			if found {
				delete(usernames, clients[websocket])
			}
			delete(clients, websocket)
			break
		}

		if clients[websocket] == "" {
			message.Name += strconv.Itoa(counter)
			_, found := usernames[message.Name]
			for {
				if !found {
					usernames[message.Name] = websocket
					break
				}
				message.Name += strconv.Itoa(counter)
				_, found = usernames[message.Name]
			}
			counter += 1
			clients[websocket] = message.Name
		}
		fmt.Println(message.Target)
		message.Message = time.Now().Format("15:04:05") + " *" + clients[websocket] + "*     " + message.Message
		broadcast <- message
	}
}

// Part 2
func broadcastMessagesToClients() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteJSON(message.Message)
			if err != nil {
				log.Printf("error occurred while writing message to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

//Part 3
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/echo", HandleClients)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting http server :: ", err)

		return
	}
}

//Final Part Over
/*
*1. generate unique ID for every client

!2. send message to selected Client any specific client

!3. Users can join multiple channels at the same time, as well as hold 1:1 chats with individual users who may or may not be in a shared channel.

*4. Each username must be unique within the current session.
*/
