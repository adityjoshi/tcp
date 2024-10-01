package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer listener.Close()

	log.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		log.Println("New client connected:", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from reverse proxy:", err)
			return
		}

		log.Printf("Received from reverse proxy: %s\n", buffer[:n])

		response := "Hello from the server!"
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Println("Error sending response to reverse proxy:", err)
		} else {
			log.Println("Response sent to reverse proxy.")
		}
	}
}
