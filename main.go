package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error connecting to reverse proxy:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Connected to the reverse proxy. Type your messages below (type 'exit' to quit):")

	for {
		fmt.Print("> ")
		message, _ := reader.ReadString('\n')
		message = message[:len(message)-1] // Remove newline character

		if message == "exit" {
			break
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message to reverse proxy:", err)
			return
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response from reverse proxy:", err)
			return
		}

		fmt.Printf("Received from reverse proxy: %s\n", buffer[:n])
	}
}
