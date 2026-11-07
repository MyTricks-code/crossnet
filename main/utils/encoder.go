package utils

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

func encode(path string, conn net.Conn) {
	file, err := os.Open(path)
	if err!=nil{
		log.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	info, err:=file.Stat()
	if err!=nil{
		panic(err)
	}
	filename := filepath.Base(path)
	size := uint64(info.Size())

	err = binary.Write(conn, binary.BigEndian, uint32(len(filename)))
	if err!=nil{
		panic(err)
	}

	_, err = conn.Write([]byte(filename))
	if err!=nil{
		panic(err)
	}

	err = binary.Write(conn, binary.BigEndian, size)
	if err!=nil{
		panic(err)
	}
	
	_, err =io.CopyN(conn, file, int64(size))
	if err != nil {
		log.Println("Error sending file:", err)
		return
	}
	log.Println("File sent successfully")
}