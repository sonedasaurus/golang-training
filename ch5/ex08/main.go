package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	node := ElementByID(doc, os.Args[2])
	if node != nil {
		fmt.Printf("idが'%s'のHTML要素を発見しました\n", os.Args[2])
		fmt.Printf("<%s", node.Data)
		for _, a := range node.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
		}
		fmt.Printf(">\n")
	} else {
		fmt.Printf("idが'%s'のHTML要素を発見できませんでした\n", os.Args[2])
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil && pre(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, id, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil && post(n, id) {
		return n
	}
	return nil
}

var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	depth++
	return false
}

func endElement(n *html.Node, id string) bool {
	depth--
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
			}
		}
	}
	return false
}
