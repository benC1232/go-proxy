package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//connecting to proxy
	address := "127.0.0.1:5577"

	conn, err := net.Dial("tcp", address)
	// check if connection was successfully established
	if err != nil {
		fmt.Println("The following error occured", err)
	}
	//reading string from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter the message you want to sent to the server: ")
	message, _ := reader.ReadString('\n')
	//sending string to proxy
	conn.Write([]byte(message))
	//closing socket
	conn.Close()
}
