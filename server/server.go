package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for {
		conn, _ := ln.Accept()
		fmt.Println("[Server] Accepted new connection.")
		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
	reader := bufio.NewReader(client)
	for {
		incomingMessage, _ := reader.ReadString('\n')
		fmt.Printf("[Server] Message received from client #%d: %s", clientid, incomingMessage)
		msgs <- Message{sender: clientid, message: incomingMessage}
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	nextClientID := 0

	//Start accepting connections
	go acceptConns(ln, conns)
	fmt.Println("[Server] Listening for connections...")

	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
			clients[nextClientID] = conn
			go handleClient(conn, nextClientID, msgs)
			fmt.Printf("[Server] Started handling new client #%d.\n", nextClientID)
			nextClientID++
		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for i := 0; i < nextClientID; i++ {
				if i != msg.sender {
					fmt.Fprint(clients[i], msg.message)
				}
			}
			fmt.Printf("[Server] Sent message from client #%d to all other clients.\n", msg.sender)
		}
	}
}
