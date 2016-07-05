package main

import (
	"fmt"
	"log"
	"os"

	"./links"
)

type Link struct {
	link  string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []Link)
	unseenLinks := make(chan Link)

	links := []Link{}
	for _, l := range os.Args[1:] {
		links = append(links, Link{l, 0})
	}
	go func() { worklist <- links }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.depth < 2 {
					fmt.Println(link.depth)
					foundLinks := crawl(link.link)
					links := []Link{}
					for _, l := range foundLinks {
						links = append(links, Link{l, link.depth + 1})
					}
					go func() { worklist <- links }()
				}
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.link] {
				seen[link.link] = true
				unseenLinks <- link
			}
		}
	}
}
