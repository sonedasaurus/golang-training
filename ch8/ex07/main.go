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

type Link struct {
	link  string
	depth int
}

func crawl(u string) []string {
	parsedUrl, err := url.Parse(u)
	if err := os.MkdirAll("./archive/"+parsedUrl.Host+parsedUrl.Path, 0777); err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("./archive/"+parsedUrl.Host+parsedUrl.Path+"/index.html", fetch.Fetch(u), os.ModePerm)
	list, err := links.Extract(u)
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
				flag := false
				for _, l := range os.Args[1:] {
					url1, err := url.Parse(l)
					if err != nil {
						log.Print(err)
						os.Exit(1)
					}
					url2, err := url.Parse(link.link)
					if err != nil {
						log.Print(err)
						os.Exit(1)
					}
					if url1.Host != url2.Host {
						flag = true
					}
				}
				if flag == true {
					fmt.Println("test")
					continue
				}
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
