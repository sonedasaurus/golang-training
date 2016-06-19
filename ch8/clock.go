package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type listenPort int
type portFlag struct{ listenPort }

func (p listenPort) String() string { return fmt.Sprintf("%d", p) }
func PortFlag(name string, value listenPort, usage string) *listenPort {
	f := portFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.listenPort
}

func (f *portFlag) Set(s string) error {
	var port int
	fmt.Sscanf(s, "%d", &port) // no error check needed
	if 0 <= port && port <= 65535 {
		f.listenPort = listenPort(port)
		return nil
	}
	return fmt.Errorf("invalid port %q", s)
}

var port = PortFlag("port", 8000, "port")

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	host := "localhost:" + strconv.Itoa(int(*port))
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
