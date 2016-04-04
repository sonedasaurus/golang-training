package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	if n.FirstChild != nil {
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}
	if n.NextSibling != nil {
		w, i := countWordsAndImages(n.NextSibling)
		words += w
		images += i
	}
	return words, images
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./ex04 http://example.com")
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Printf("words = %d\n", words)
		fmt.Printf("images = %d\n", images)
	}
}
