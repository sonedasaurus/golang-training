package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"./fetch"
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
	if err := os.MkdirAll("./archive/"+parsedUrl.Host+parsedUrl.Path, 0777); err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("./archive/"+parsedUrl.Host+parsedUrl.Path+"/archive.html", fetch.Fetch(u), os.ModePerm)
	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
