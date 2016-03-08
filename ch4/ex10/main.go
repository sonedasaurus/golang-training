package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

//!+
func main() {
	t := time.Now()
	beforeMonth := t.AddDate(0, -1, 0)
	beforeYear := t.AddDate(0, -1, 0)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		// fmt.Printf("#%-5d %9.9s %.55s\n",
		// item.Number, item.User.Login, item.Title)
		fmt.Println(item.CreatedAt)
	}
}
