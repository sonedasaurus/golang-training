package main

import (
	"fmt"
	"time"
)

func pingpong() {
	var in = make(chan string)
	var out = make(chan string)
	count := 0

	go func() {
		for {
			msg := <-in
			fmt.Println(msg)
			count++
			out <- "           pong!"
		}
	}()

	go func() {
		for {
			msg := <-out
			fmt.Println(msg)
			count++
			in <- "ping!           "
		}
	}()
	in <- "ping!           "

	tick := time.Tick(1 * time.Second)
	<-tick
	fmt.Printf("count: %d\n", count)
	return
}

func main() {
	pingpong()
}
