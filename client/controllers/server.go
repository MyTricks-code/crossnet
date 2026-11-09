package controllers

import (
	"context"
	"fmt"
	"io"
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

func ServerListener() net.Listener{
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Print("Error while listening: ", err)
		return nil
	}
	fmt.Println("Server listening on 9000")
	defer listener.Close()
	return listener
}

func SpawnConnection() net.Conn{
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Print("Error starting server: ", err)
	}
	defer conn.Close()
	return conn
}

func HandleConnection(conn net.Conn, ctx context.Context) error{
	fmt.Print("looking for connections...")
	// ctx, cancel:=context.WithTimeout(ctx, 2*time.Second)
	// defer cancel()
	typebuf:=make([]byte,1)
	_,err:=io.ReadFull(conn, typebuf)
	if err!=nil{
		fmt.Println("Error reading typebuf: ",err)
		return nil
	}
	msgType:=typebuf[0]
	if msgType==HELLO {
		fmt.Println("HELLO")
	}else if msgType==2{
		fmt.Println("FILE_OFFER")
	}else if msgType==3{
		fmt.Println("ACCEPT")
	}else{
		fmt.Println("INVALID PROTOCOL")
	}
	return nil;
}