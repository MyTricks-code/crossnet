package controllers

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"io"
)

func FileSend(conn net.Conn) error{
	/*
	Order of sending: 
	File name size
	File name
	File size
	File actual bites
	*/

	var fileLocation string
	fmt.Println("Enter filePath: ")

	fmt.Scan(&fileLocation)

	file,err:=os.Open(fileLocation)
	if err!=nil{
		fmt.Println("Error opening file: ", err)
		return  err
	}
	defer file.Close()
	info , err:=file.Stat()
	if err!=nil{
		panic(err)
	}
	fileSize:=uint64(info.Size())

	filename:=filepath.Base(fileLocation)
	
	var fileNameSize uint32

	fileNameSize=uint32(len(filename))
	err=binary.Write(conn, binary.BigEndian, fileNameSize)
	if err != nil{
		fmt.Println("Error sending fileLength: ", err)
		return  err
	}
	_, err = conn.Write([]byte(filename))
	if err != nil{
		fmt.Println("Error sending fileLength: ", err)
		return  err
	}
	err=binary.Write(conn, binary.BigEndian, fileSize)
	if err != nil{
		fmt.Println("Error sending file size: ", err)
		return  err
	}
	_, err = io.CopyN(conn, file, int64(fileSize))
	if err != nil {
		fmt.Println("Error sending file:", err)
		return err
	}
	fmt.Println("File send successfully")
	return nil
}

func FileReceive(conn net.Conn) error{
	// receiving file
	err := os.MkdirAll("uploads", 0755)
	if err != nil {
		fmt.Println("Error creating uploads directory:", err)
		return err
	}

	var fileNameSize uint32
	var fileSize uint64
	err=binary.Read(conn, binary.BigEndian, &fileNameSize)
	
	fileNamebuf:=make([]byte, fileNameSize)
	_,err=io.ReadFull(conn, fileNamebuf)

	out, err := os.Create("uploads/" + string(fileNamebuf))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer out.Close()
	
	err = binary.Read(conn, binary.BigEndian, &fileSize)
	if err!=nil{
		fmt.Println("Error reading file size: ", err)
		return err
	}
	_,err=io.CopyN(out,conn,int64(fileSize))
	if err!=nil{
		fmt.Println("Error saving file: ", err)
		return err
	}
	fmt.Println("File saved")
	conn.Close()
	return nil
}