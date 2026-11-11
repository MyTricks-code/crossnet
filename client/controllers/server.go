package controllers

import (
	"fmt"
	"net"
	// "time"
)

const (
	HELLO byte = iota + 1
	FILE_OFFER
	ACCEPT
	REJECT
	FILE_DATA
	TRANSFER_COMPLETE
	TRANSFER_ERROR
)

func ServerListener() net.Listener {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Print("Error while listening: ", err)
		return nil
	}
	fmt.Println("Server listening on 9000")
	return listener
}

func SpawnConnection() net.Conn {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Print("Error starting server: ", err)
	}
	return conn
}

func HandleConnection(choice int) error {
	/*
	1. Receive
	2. Send
	*/
	if choice==1 {
		conn:=SpawnConnection()
		FileReceive(conn)
		defer conn.Close()

	}else if choice==2{
		listener:=ServerListener()
		defer listener.Close()
		for {
			// fmt.Print("Waiting for connections....")
			conn, err:=listener.Accept()
			if err!=nil {
				fmt.Println("Error stablishing a connection: ", err)
				continue
			}
			go func (conn net.Conn)  {
				FileSend(conn)
				defer conn.Close()
			}(conn)
		}

	}
	return nil
}
