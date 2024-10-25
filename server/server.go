package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

var (
	clients = make(map[net.Conn]string)
	mu      sync.Mutex
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var name string

	// Receive client name
	err := json.NewDecoder(conn).Decode(&name)
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	mu.Lock()
	clients[conn] = name
	mu.Unlock()

	fmt.Println("Client connected:", name)

	for {
		// Wait for client messages
		var msg string
		err := json.NewDecoder(conn).Decode(&msg)
		if err != nil {
			break
		}
		fmt.Printf("%s: %s\n", name, msg)
	}

	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
