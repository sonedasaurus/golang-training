package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func htmlTagCount(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	if n.FirstChild != nil {
		htmlTagCount(counts, n.FirstChild)
	}
	if n.NextSibling != nil {
		htmlTagCount(counts, n.NextSibling)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./ex01 http://example.com")
		os.Exit(1)
	}
	var counts map[string]int
	counts = map[string]int{}
	for _, url := range os.Args[1:] {
		err := findLinks(url, counts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
			continue
		}
		for key, value := range counts {
			fmt.Printf("tagName = %s, tagCount = %d\n", key, value)
		}
	}
}

func findLinks(url string, counts map[string]int) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	htmlTagCount(counts, doc)
	return nil
}
