package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"./links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(u string) []string {
	parsedUrl, err := url.Parse(u)
	if err := os.Mkdir("./archive/"+parsedUrl.Host, 0777); err != nil {
		fmt.Println(err)
	}
	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
