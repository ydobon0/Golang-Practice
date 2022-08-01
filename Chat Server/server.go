package main

//  http://localhost:8080/
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

//var channelInx int = 0
var counter int = 0
var clients = make(map[*websocket.Conn]string)
var usernames = make(map[string]*websocket.Conn)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

var channelSli = make([]ChannelMap, 0)

type Message struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Target  string `json:"target"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

type ChannelMap struct {
	chl_clients map[*websocket.Conn]string
	chl_names   map[string]*websocket.Conn
}

// {Part 1 }
func HandleClients(w http.ResponseWriter, r *http.Request) {
	//var funcID int = counter
	go broadcastMessagesToClients()

	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket :: ", err)
	}
	defer websocket.Close()
	clients[websocket] = ""

	//log.Printf("*** HandleClients%d 1", funcID)
	for {
		var message Message

		err := websocket.ReadJSON(&message)
		//log.Printf("*** HandleClients%d 2", funcID)

		if err != nil {
			log.Printf("error occurred while reading message : %v", err)
			deleteUser(websocket)
			updateUserList()
			break
		}
		// implement unique name and then add the unique name into maps of clients, usernames
		if clients[websocket] == "" && message.Name != "" {
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

			if channelSli[0].chl_clients[websocket] == "" { // add user and websocket to channel 1
				channelSli[0].chl_clients[websocket] = message.Name
				channelSli[0].chl_names[message.Name] = websocket
			}

			updateUserList()
		}

		message.Name = clients[websocket] //set the sender's name to the correct value

		//command starts with / and ends with ::
		if message.Type == "message" && message.Message[0:1] == "/" && strings.Contains(message.Message, "::") {
			processCommand(&message)
		}

		// inx, _ := strconv.Atoi(message.Channel)
		// if channelSli[inx].chl_clients[websocket] == "" && message.Name != "" { // add user and websocket to channel slices
		// 	channelSli[inx].chl_clients[websocket] = message.Name
		// 	channelSli[inx].chl_names[message.Name] = websocket
		//}

		if message.Type == "message" {
			message.Message = time.Now().Format("15:04:05") + " *" + clients[websocket] + "*     " + message.Message
			broadcast <- message
		} else if message.Type == "channel" {
			addChannel()
			updateChannel()
		} else if message.Type == "update" {
			updateUserList()
			updateChannel()
		}
	}
}

// Part 2
func broadcastMessagesToClients() {
	//var funcID int = counter
	//counter++
	//log.Printf("+++ broadcastMessagesToClients%d 1", funcID)
	for {
		message := <-broadcast
		//log.Printf("+++ broadcastMessagesToClients%d 2", funcID)

		if message.Type == "message" {
			_, found := usernames[message.Target] //check if the target is a valid user
			if found {                            //Send message only to the target and the sender
				usernames[message.Name].WriteJSON(message)
				if message.Name != message.Target {
					usernames[message.Target].WriteJSON(message)
				}
			} else if message.Channel == "0" { // send  message to ALL connected clients
				for client := range clients {
					err := client.WriteJSON(message)
					if err != nil {
						log.Printf("error occurred while writing message to client: %v", err)
						client.Close()
						delete(clients, client)
					}
				}
			} else { // send message to clients in a channel
				inx, _ := strconv.Atoi(message.Channel)
				_, found := channelSli[inx].chl_names[message.Name] // to make sure the sender is in the channel
				if found {
					for client := range channelSli[inx].chl_clients {
						err := client.WriteJSON(message) //.Message)
						if err != nil {
							log.Printf("error occurred while writing message to client: %v", err)
							client.Close()
							delete(clients, client)
						}
					}
				} else {
					message.Message = time.Now().Format("15:04:05") + " *" + message.Name + "*     You are not in channel " + strconv.Itoa(inx+1)
					err := usernames[message.Name].WriteJSON(message) // send the message to the sender
					if err != nil {
						log.Printf("error occurred while writing message to client: %v", err)
					}
				}
			}
		}
	}
}

//Part 3
func main() {
	//log.Printf("### main")
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

//Part 4: Alan code
func init() { // called before main()
	addChannel() // add the verfirst channel
}

// notify all connected clients that to update "To" drop down list
func updateUserList() {
	var message Message
	message.Message = "ALL"
	message.Type = "user"

	for user := range usernames {
		message.Message += "," + user
	}

	for client := range clients {
		err := client.WriteJSON(message) //.Message)
		if err != nil {
			log.Printf("updateList: error occurred while writing message to client: %v", err)
			client.Close()
			deleteUser(client)
		}
	}
}

// delete the user from maps, include from all maps in channel slices
func deleteUser(wskt *websocket.Conn) {
	for ii := 0; ii < len(channelSli); ii++ {
		name, exist := channelSli[ii].chl_clients[wskt]
		if exist {
			_, find := channelSli[ii].chl_names[name]
			if find {
				delete(channelSli[ii].chl_names, name)
			}
			delete(channelSli[ii].chl_clients, wskt)
		}
	}

	_, found := usernames[clients[wskt]]
	if found {
		delete(usernames, clients[wskt])
	}
	delete(clients, wskt)

}

// add a ChannelMap to "channelSli"
func addChannel() {
	chl := new(ChannelMap)
	chl.chl_clients = make(map[*websocket.Conn]string)
	chl.chl_names = make(map[string]*websocket.Conn)

	channelSli = append(channelSli, *chl)
}

// notify all connected clients that to update "Channel" drop down list
func updateChannel() {
	var message Message
	message.Channel = "0"
	message.Type = "channel"

	for ii := 1; ii < len(channelSli); ii++ {
		message.Channel += "," + strconv.Itoa(ii)
	}

	for client := range clients {
		err := client.WriteJSON(message) //.Message)
		if err != nil {
			log.Printf("updateChannel: error occurred while writing message to client: %v", err)
			client.Close()
			deleteUser(client)
		}
	}
}

func processCommand(message *Message) {
	str := strings.Split(message.Message, "::")
	fmt.Println(str)
	message.Message = str[1]
	strTemp := strings.ToLower(str[0][1:])
	strTemp = strings.TrimSpace(strTemp)
	fmt.Println(strTemp)
	command := strings.Split(strTemp, " ")

	switch command[0] {
	case "time": //command for getting the current time
		message.Message = time.Now().Format("15:04:05")

	case "join": //command for join a channel
		strTemp = strings.TrimSpace(command[1])
		chl := strings.Split(strTemp, "channel")
		nn, _ := strconv.Atoi(chl[1])
		if nn > 0 && nn <= len(channelSli) {
			//message.Channel = strconv.Itoa(nn - 1)
			nn--
			websocket, find := usernames[message.Name]
			if find && channelSli[nn].chl_clients[websocket] == "" { // add user and websocket to channel slices
				channelSli[nn].chl_clients[websocket] = message.Name
				channelSli[nn].chl_names[message.Name] = websocket
			}
			message.Message = message.Name + " joined " + command[1]
		} else {
			message.Message = command[1] + " is not exist"
			message.Target = message.Name
		}

	case "leave":
		strTemp = strings.TrimSpace(command[1])
		chl := strings.Split(strTemp, "channel")
		nn, _ := strconv.Atoi(chl[1])
		if nn > 0 && nn <= len(channelSli) {
			ii := nn - 1
			websocket, find := channelSli[ii].chl_names[message.Name]
			if find {
				name, exist := channelSli[ii].chl_clients[websocket]
				if exist {
					delete(channelSli[ii].chl_names, name)
				}
				delete(channelSli[ii].chl_clients, websocket)
				message.Message = message.Name + " left " + command[1]
			} else {
				message.Message = message.Name + " is not in " + command[1]
			}
		} else {
			message.Message = command[1] + " is not exist"
			message.Target = message.Name
		}

	default: //command for sending a private message to another user
		user := strings.TrimSpace(str[0][1:])
		_, found := usernames[user]
		if found {
			message.Target = user
		}
	}
}

//Final Part Over
/*
1. generate unique ID for every client

2. send message to selected Client any specific client

3. Users can join multiple channels at the same time, as well as hold 1:1 chats with individual users who may or may not be in a shared channel.

4. Each username must be unique within the current session.
*/
