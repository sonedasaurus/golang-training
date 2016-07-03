package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		println("error tcp resolve failed", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		conn.CloseRead()
		log.Println("close read")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	log.Println("close write")
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
