package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client struct {
	id  string
	msg chan string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.msg <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.msg <- "-----participants------"
			for client := range clients {
				cli.msg <- client.id
			}
			cli.msg <- "-----------------------"

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := client{}
	ch.msg = make(chan string) // outgoing client messages
	go clientWriter(conn, ch.msg)

	who := conn.RemoteAddr().String()
	ch.id = who
	ch.msg <- "You are" + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
