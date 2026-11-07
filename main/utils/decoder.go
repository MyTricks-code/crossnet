package utils

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
)

func Decoder(conn net.Conn){
	err := os.MkdirAll("uploads", 0755)
	if err != nil {
		log.Println("Error creating uploads directory:", err)
		return
	}

	var nameLen uint32
	err = binary.Read(conn, binary.BigEndian, &nameLen)
	if err!=nil{
		log.Println("Error reading name length: ", err)
		return
	}
	namebuf := make([]byte, nameLen)
	_, err =io.ReadFull(conn, namebuf)
	if err!=nil{
		log.Println("Error reading file name: ", err)
		return
	}

	out, err := os.Create("uploads/" + string(namebuf))
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	var fileSize uint64
	err = binary.Read(conn, binary.BigEndian, &fileSize)
	if err!=nil{
		log.Println("Error reading file size: ", err)
		return
	}

	err = CopyFile(conn, out, fileSize)
	if err!=nil{
		log.Println("Error reading file content: ", err)
		return
	}
	log.Println("[Decoder] File recieved and created")
}