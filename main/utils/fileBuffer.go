package utils

import (
	"io"
	"net"
	"os"
)

func CopyFile(conn net.Conn, file *os.File, size uint64) error{
	const chunkSize = 32*1024;
	buf:=make([]byte, chunkSize)

	var sent uint64; 
	for sent<size{
		want:=uint64(len(buf))
		remaining:=size-sent;
		if remaining<want{
			want=remaining
		}
		n,err:=io.ReadFull(conn, buf[:want])
		if err != nil {
            return err
        }
		written:=0
		for written<n{
			m,err:=file.Write(buf[written:n])
			if err != nil {
                return err
            }
            written += m
		}
		sent += uint64(n)
	}
	return nil;
}