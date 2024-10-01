package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Server is listening on port 8081")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleClients(conn)
	}
}

func handleClients(conn net.Conn) {
	defer conn.Close()

	serverConn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to backend server:", err)
		return
	}
	defer serverConn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		fmt.Printf("Received: %s\n", buffer[:n])

		_, err = serverConn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error sending data to backend:", err)
			return
		}

		responseBuffer := make([]byte, 1024)
		responseBytes, err := serverConn.Read(responseBuffer)
		if err != nil {
			fmt.Println("Error reading response from backend:", err)
			return
		}

		_, err = conn.Write(responseBuffer[:responseBytes])
		if err != nil {
			fmt.Println("Error sending response to client:", err)
			return
		}
	}
}
