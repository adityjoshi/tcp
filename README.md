# TCP Server and Reverse Proxy Project

## Overview

This project involves the creation of a TCP server and a reverse proxy from scratch. It aims to provide insights into how a TCP server handles multiple concurrent requests while routing client traffic efficiently through a reverse proxy.

## Key Features

- **TCP Server**: Built to handle multiple simultaneous client connections.
- **Reverse Proxy**: Routes requests from clients to the backend server and returns the responses.
- **Concurrency Management**: Utilizes Go's goroutines for efficient handling of concurrent connections.
- **Real-Time Communication**: Enables seamless data transfer between clients and servers.

## Technical Details

### TCP Server

- **Implementation**: The TCP server listens for incoming connections and processes requests concurrently.
- **Concurrency**: Each client connection is handled in a separate goroutine to ensure non-blocking operations.
- **Data Handling**: The server reads data from clients, processes it, and sends responses back.

### Reverse Proxy

- **Routing Logic**: The reverse proxy accepts client requests and forwards them to the TCP server.
- **Connection Management**: It establishes connections to the backend server and handles responses efficiently.
- **Error Handling**: Implements robust error handling to manage failed connections and data transfer issues.

## How to Run

1. **Run the TCP Server**:
   ```go
   go run server/server.go
2. **Run the reverse proxy server**:
   ```go
   go run reverse_proxy/reverseProxy.go
3. **Run the main.go**:
   ```go
   go run main.go

4. **Send Requests using netcat: Open a terminal and use the following command to send a request to the reverse proxy**:
   ```bash
   echo "Hello TCP!" | nc localhost 8081
