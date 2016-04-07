package main

import (
	"errors"
	"fmt"
)

func main() {
	err := panicAndRecover()
	fmt.Println(err)
}

func panicAndRecover() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New("recoverd!")
		}
	}()
	panic("hogehoge")
}
