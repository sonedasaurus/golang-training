package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	nodes := getTag(nil, doc, name)
	return nodes
}

func getTag(result []*html.Node, n *html.Node, name []string) []*html.Node {
	if n.Type == html.ElementNode {
		for _, value := range name {
			if value == n.Data {
				result = append(result, n)
			}
		}
	}
	if n.FirstChild != nil {
		result = getTag(result, n.FirstChild, name)
	}
	if n.NextSibling != nil {
		result = getTag(result, n.NextSibling, name)
	}
	return result
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./ex02 http://example.com")
		os.Exit(1)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println("Get Error")
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	nodes := ElementsByTagName(doc, os.Args[2:]...)
	for i, value := range nodes {
		fmt.Printf("%d: Data = %s \n", i, value.Data)
	}
}
