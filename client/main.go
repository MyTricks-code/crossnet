package main

import (
	"context"
	"crossnetClient/controllers"
	"fmt"
	"net"
)

func main() {
	fmt.Println("welcome to Crossnet")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Print("Error starting server: ", err)
	}
	defer conn.Close()

	fmt.Print("Server running on: localhost:8080")
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		fmt.Print("Error while listening: ", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("Error while connecting: ", err)
			break
		}
		defer conn.Close()
		ctx:=context.Background()
		controllers.HandleConnection(conn, ctx)
	}

}
