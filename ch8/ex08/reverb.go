package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	timeout := 10
	input := bufio.NewScanner(c)
	tick := time.Tick(1 * time.Second)
	inputch := make(chan struct{})
	go func() {
		for input.Scan() {
			inputch <- struct{}{}
		}
	}()
	for count := 0; count < timeout; count++ {
		select {
		case <-inputch:
			count = 0
			go echo(c, input.Text(), 1*time.Second)
		case <-tick:
			// do nothing
		}
	}
	c.Close()
	return
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
