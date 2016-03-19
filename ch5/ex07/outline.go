package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout // modified during testing

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild == nil {
		fmt.Fprintf(out, "%*s<%s/", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(out, " %s=\"%s\"", a.Key, a.Val)
		}
		fmt.Fprintf(out, ">\n")
		depth++
		return
	}
	if n.Type == html.ElementNode {
		fmt.Fprintf(out, "%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(out, " %s=\"%s\"", a.Key, a.Val)
		}
		fmt.Fprintf(out, ">\n")
		depth++
		return
	}
	if n.Type == html.TextNode {
		fmt.Fprintf(out, "%*s%s\n", depth*2, "", n.Data)
		return
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild == nil {
		depth--
		return
	}
	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
		return
	}
}
