package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func handleConn(r string, c net.Conn) {
	defer c.Close()
	fmt.Println(r)
	mustCopy(os.Stdout, c)
}

func main() {
	regions := []string{"NewYork", "Tokyo", "London"}
	for _, region := range regions {
		if os.Getenv(region) != "" {
			conn, err := net.Dial("tcp", os.Getenv(region))
			if err != nil {
				log.Fatal(err)
			}
			go handleConn(region, conn)
		}
	}
	for { // loop
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
