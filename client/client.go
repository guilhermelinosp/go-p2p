package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send client name to server
	err = json.NewEncoder(conn).Encode(name)
	if err != nil {
		fmt.Println("Error sending name:", err)
		return
	}

	// Listen for messages from the server
	go func() {
		for {
			var msg string
			err := json.NewDecoder(conn).Decode(&msg)
			if err != nil {
				break
			}
			fmt.Println(msg)
		}
	}()

	// Send messages to server
	for {
		var msg string
		fmt.Print("Enter message: ")
		fmt.Scanln(&msg)
		err := json.NewEncoder(conn).Encode(msg)
		if err != nil {
			break
		}
	}
}
