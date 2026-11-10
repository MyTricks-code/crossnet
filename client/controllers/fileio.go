package controllers

import (
	"encoding/binary"
	"fmt"
	"net"
	"path/filepath"
)

func FileSend(conn net.Conn) error{
	var fileLocation string
	fmt.Println("Enter filePath: ")
	fmt.Scan(&fileLocation)
	filename:=filepath.Base(fileLocation)
	
	var fileNameSize uint8
	fileNameSize=uint8(len(filename))
	err:=binary.Write(conn, binary.BigEndian, fileNameSize)

	if err != nil{
		fmt.Println("Error sending fileLength: ", err)
		return  err
	}
	_, err = conn.Write([]byte(filename))
	if err != nil{
		fmt.Println("Error sending fileLength: ", err)
		return  err
	}
	return nil
}

func FileOffer(conn net.Conn) error{
	
}