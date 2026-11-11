package main

import (
	"crossnetClient/controllers"
	"fmt"
	// "net"
)

func main() {
	fmt.Println("welcome to Crossnet")

	menu := "1. Receive\n 2. Send"
	fmt.Println("Menu: \n", menu)

	var choice uint8
	fmt.Println("Enter choice: ")
	fmt.Scan(&choice)

	switch choice {	
	case 1:
		controllers.HandleConnection(1)
		fmt.Println("Receive selected")
	case 2:
		fmt.Println("Send selected")
		controllers.HandleConnection(2)
	case 3:
		fmt.Println("Closing connections")
		controllers.HandleConnection(3)
	default:
		fmt.Println("Invalid choice")
	}
}
