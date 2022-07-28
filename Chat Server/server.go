package main

//  http://localhost:8080/

/*
First Create a Audit Report of the project like flow of application step by step then only you can add more features in Right Place
*/

import (
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
	Type    string `json:"type"`
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
			updateList()
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
			updateList()
		}
		//fmt.Println(message.Target)
		message.Name = clients[websocket] //set the sender's name to the correct value
		if message.Type == "message" {
			message.Message = time.Now().Format("15:04:05") + " *" + clients[websocket] + "*     " + message.Message
			broadcast <- message
		} //else {
		// 	message.Message = "ALL"
		// 	for user := range usernames {
		// 		message.Message += "," + user
		// 	}
		// 	websocket.WriteJSON(message)
		// }
	}
}

// Part 2
func broadcastMessagesToClients() {
	for {
		message := <-broadcast

		if message.Type == "message" {
			_, found := usernames[message.Target] //check if the target is a valid user
			if found {                            //if it is send the message only to the target and the sender
				usernames[message.Target].WriteJSON(message) //.Message
				usernames[message.Name].WriteJSON(message)   //.Message
			} else {
				for client := range clients {
					err := client.WriteJSON(message) //.Message)
					if err != nil {
						log.Printf("error occurred while writing message to client: %v", err)
						client.Close()
						delete(clients, client)
					}
				}
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

func updateList() {
	var message Message
	message.Message = "ALL"
	message.Type = "require"
	for user := range usernames {
		message.Message += "," + user
	}
	for client := range clients {
		err := client.WriteJSON(message) //.Message)
		if err != nil {
			log.Printf("updateList: error occurred while writing message to client: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
	//websocket.WriteJSON(message)

}

/*
Users can run commands using the same interface as a chat, but commands start with /. These commands allow the user to join a channel, leave a channel, start a direct chat with another user, get the current time, or even interact with a channel's bot.
*/
