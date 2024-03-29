package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port 8000")
	conn, err := ln.Accept()

	if err != nil {
		log.Fatal(err)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Message Received:", string(message))
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))
	}
}
