package controllers

import (
	"context"
	"fmt"
	"net"
	"time"
)

func HandleConnection(conn net.Conn, ctx context.Context) error{
	fmt.Print("looking for connections...")
	// ctx, cancel:=context.WithTimeout(ctx, 2*time.Second)
	// defer cancel()

	
}