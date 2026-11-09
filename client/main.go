package main

import (
	"context"
	"crossnetClient/controllers"
	"fmt"
	// "net"
)

func main() {
	fmt.Println("welcome to Crossnet")

	menu:="1. Receive\n 2.Send"
	fmt.Println("Menu: \n", menu)

	var choice uint8
	fmt.Println("Enter choice: ")
	fmt.Scan(&choice)

	if choice==1{
		listener:=controllers.ServerListener()
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
	}else if choice==2{
		conn:=controllers.SpawnConnection()
		_,err:=conn.Write([]byte{1})
		if err!=nil{
			fmt.Println("Error writing to connection: ", err)
		}

	}

	

}
