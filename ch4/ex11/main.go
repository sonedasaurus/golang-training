package main

import (
	"fmt"
	"os"

	"./github"
)

func main() {
	err := github.CreateIssues(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
}
