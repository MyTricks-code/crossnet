package server

import (
	"log"
	"net"
)

func Start() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err!=nil {
		log.Fatal("[Error starting server]: ",err)
	}
	defer listener.Close()

	conn,err:=listener.Accept()
	if err!=nil {
		log.Println("[server] Connections missed: ",err)
	}

	defer conn.Close()
}