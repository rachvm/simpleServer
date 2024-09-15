package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server using a tcp connection to localHost:8080
	conn, err := net.Dial("tcp", "localHost:8080")
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}
	defer conn.Close()

	// Manually construct the HTTP POST request
	request := "POST / HTTP/1.1\r\n" +
		"Host: localhost:8080\r\n" +
		"Content-Type: application/json\r\n" +
		"Content-Length: 18\r\n" + // Adjust content length based on the JSON data
		"\r\n" +
		`{"name":"golang"}` + "\r\n"

	// Send the request to the server
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}

	// Read response from the server
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response: ", err)
	} 

	fmt.Println("Response from server: ", string(buffer[:n]))
}