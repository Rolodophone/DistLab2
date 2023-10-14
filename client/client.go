package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"syscall"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		incomingMessage, err := reader.ReadString('\n')
		if err == nil {
			fmt.Print(incomingMessage)
		} else {
			fmt.Printf("[Client] Error reading from server:\n%s\n", err)
		}
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("")
		scanner.Scan()
		_, err := fmt.Fprintln(conn, scanner.Text())
		if err != nil {
			fmt.Printf("[Client] Error sending message to server:\n%s\n", err)
		}
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	//TODO Try to connect to the server
	fmt.Printf("[Client] Trying to connect to server at %s.\n", *addrPtr)
	conn, err := net.Dial("tcp", *addrPtr)
	if err == nil {
		fmt.Println("[Client] Connected to the server.")
	} else {
		fmt.Printf("[Client] Error connecting to server. Perhaps the IP address is wrong?\n%s\n", err)
		syscall.Exit(1)
	}

	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	go write(conn)

	<-make(chan int) //block main to keep read and write goroutines alive
}
