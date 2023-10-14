package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		incomingMessage, _ := reader.ReadString('\n')
		fmt.Print(incomingMessage)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("")
		scanner.Scan()
		fmt.Fprintln(conn, scanner.Text())
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	//TODO Try to connect to the server
	conn, _ := net.Dial("tcp", *addrPtr)
	fmt.Println("[Client] Dialled the server.")

	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	go write(conn)

	<-make(chan int) //block main to keep read and write goroutines alive

	//TODO Start getting and sending user messages.
}
