package controllers

import (
	"fmt"
	"net"
	"sync"
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

var globalListener *net.Listener

func HandleConnection(choice int) error {
	/*
		1. Receive
		2. Send
	*/
	if choice == 1 {
		conn := SpawnConnection()
		defer conn.Close()
		FileReceive(conn)

	} else if choice == 2 {
		var wg sync.WaitGroup
		listener := ServerListener()
		globalListener = &listener
		defer listener.Close()
		
		var fileLocation string
		fmt.Println("Enter filePath: ")
		fmt.Scan(&fileLocation)

		for {
			// fmt.Print("Waiting for connections....")
			
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error stablishing a connection: ", err)
				break
			}
			wg.Add(1)
			go func(conn net.Conn, fileLocation string) {
				defer wg.Done()
				defer conn.Close()
				FileSend(conn, fileLocation)
			}(conn, fileLocation)
		}
		wg.Wait()
		fmt.Println("File send successfully. Close the terminal")
		

	} else if choice == 3 {
		gl := *globalListener
		gl.Close()
	}
	return nil
}
