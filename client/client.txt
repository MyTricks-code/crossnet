package main

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
	if err != nil {
		log.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	filename := filepath.Base(path)
	size := uint64(info.Size())

	err = binary.Write(conn, binary.BigEndian, uint32(len(filename)))
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte(filename))
	if err != nil {
		panic(err)
	}

	err = binary.Write(conn, binary.BigEndian, size)
	if err != nil {
		panic(err)
	}

	_, err = io.CopyN(conn, file, int64(size))
	if err != nil {
		log.Println("Error sending file:", err)
		return
	}
	log.Println("File sent successfully")
}

func main() {
    dev := os.Getenv("DEV")

    addr := "localhost:8080"
    path := "./test.jpg"

    if dev == "docker" {
        addr = "server:8080"
        path = "/app/client/test.jpg"
    }

    conn, err := net.Dial("tcp", addr)
    if err != nil {
        log.Println("[Client] Error establishing connection:", err)
        return
    }
    defer conn.Close()

    encode(path, conn)
    conn.(*net.TCPConn).CloseWrite()

    buff := make([]byte, 1024)
    n, err := conn.Read(buff)
    if err != nil {
        log.Println("Read Error:", err)
        return
    }

    log.Println("Server responded:", string(buff[:n]))
}
