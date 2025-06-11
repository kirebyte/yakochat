package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader to convert HTTP to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// List of connected clients
var clients = make(map[*websocket.Conn]bool)

// Channel to broadcast messages to all clients
var broadcast = make(chan string)

// WebSocket handler
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Add the client to the list of connected clients
	clients[ws] = true

	for {
		var msg string
		// Read messages from the client
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, ws)
			break
		}
		// Send the message to the broadcast channel
		broadcast <- msg
		fmt.Println(msg)
	}
}

// Function to send messages to all connected clients
func handleMessages() {
	for {
		// Read messages from the broadcast channel
		msg := <-broadcast
		// Send the message to all connected clients
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// Function to read from the console and send messages to clients
func consoleReader() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		// Send the console message to the broadcast channel
		broadcast <- msg
	}
}

func main() {
	// Route to serve the chat page
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// WebSocket route
	http.HandleFunc("/ws", handleConnections)

	// Start reading messages from the console
	go consoleReader()

	// Start handling messages between clients and server
	go handleMessages()

	// Start the server on port 8080
	fmt.Println("Server listening on http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
