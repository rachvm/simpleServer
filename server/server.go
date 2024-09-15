package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Socket Creation - creating a connection between a client and a server over tcp
	// TCP transmission control protocol - communications standard that enables application programs and computing deviced to exchange messages over a network

	// Listen creates the server
	l, err := net.Listen("tcp", "localHost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// defer - delays the execution of a function until the surrounding function finishes
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// If err == nill then run the connection
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// Handle the connection in a new go routine.
	// The loop then returns accepting, so that multiple connections may be served concurrentl.
	// Don't do conn.Close() (shutdown the connection) till this function has finished
	defer conn.Close()

	// Read the connection
	// buffer is a temporary memory storage area used to hold datd while it is transfered from one place to another
	// bufio is the go package which includes buffered I/O functions that wrap around unbuffered readers and writers to make reading and writing more efficent
	// I/O operations - input/output operations refer tto processes by which a computer system exchanges data with external devices
	// buffer.NewReader - wraps the connection in a buffer so that data can be read in efficent chunks
	reader := bufio.NewReader(conn)
	// Readstring('\n')  Reads from the buffer until it encounters a newline character - more efficent that reading byte to byte
	req, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request: ", err)
		return
	}

	fmt.Println("Received request: ", req)

	// Process request 
	// Respond to client
	response := "HTTP/1.1 200 OK\r\nContent-Type: test/plain\r\n\r\nHello, world!\n"
	conn.Write([]byte(response))
}
