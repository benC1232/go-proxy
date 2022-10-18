package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:5578")
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
	conn.Close()
	listener.Close()
}
