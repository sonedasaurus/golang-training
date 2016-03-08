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
	beforeYear := t.AddDate(-1, 0, 0)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("一ヶ月未満")
	for _, item := range result.Items {
		if item.CreatedAt.After(beforeMonth) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("一年未満")
	for _, item := range result.Items {
		if item.CreatedAt.After(beforeYear) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("一年以上")
	for _, item := range result.Items {
		if item.CreatedAt.Before(beforeYear) {
			fmt.Printf("#%-5d %9.9s %.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
	}
}
