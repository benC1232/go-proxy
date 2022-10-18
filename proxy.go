package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//listening to client
	listener, err := net.Listen("tcp", "127.0.0.1:5577")
	if err != nil {
		fmt.Println("error listening: " + err.Error())
		os.Exit(1)
	}
	conn, err := listener.Accept()
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("error reading: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("message received:" + string(buffer))
	//connecting to server
	address := "127.0.0.1:5578"
	conn, err = net.Dial("tcp", address)
	// check if connection was successfully established
	if err != nil {
		fmt.Println("The following error occurred", err)
	}
	//sending string to server
	conn.Write([]byte(buffer))
	//closing socket
	conn.Close()
	listener.Close()
}
