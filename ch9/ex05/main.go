package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func pipeline(size int) (chan<- struct{}, <-chan struct{}) {
	var in = make(chan struct{})
	var out = make(chan struct{})

	pipe := make([]chan struct{}, size)
	for i, _ := range pipe {
		pipe[i] = make(chan struct{})
	}

	go func() {
		<-in
		pipe[0] <- struct{}{}
	}()
	for i := 0; i < size-1; i++ {
		idx := i
		go func() {
			<-pipe[idx]
			pipe[idx+1] <- struct{}{}
		}()
	}
	go func() {
		<-pipe[size-1]
		out <- struct{}{}
	}()

	return in, out
}

func main() {
	usage := "Usage: ./ex05 100"
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(usage)
		os.Exit(1)
	}

	fmt.Println("now pipeline creating ...")
	in, out := pipeline(num)
	fmt.Println("start")
	in <- struct{}{}
	start := time.Now()
	<-out
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
