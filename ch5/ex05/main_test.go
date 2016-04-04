package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args string
		want map[string]int
	}{
		{`<html><head><style src="/path/style"></style></head><body><a href="foo">Foo Bar Hoge</a></body></html>`, map[string]int{"words": 3, "images": 0}},
		{`<html><head><style src="/path/style"></style></head><body><a href="foo">Foo Bar</a></body></html>`, map[string]int{"words": 2, "images": 0}},
		{`<html><head><script src="/path/script"></script></head><body><a href="foo">Foo</a></body></html>`, map[string]int{"words": 1, "images": 0}},
		{`<html><head><img src="/path/img"></img></head><body><a href="foo">Foo</a></body></html>`, map[string]int{"words": 1, "images": 1}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("(%q)", test.args)
		doc, err := html.Parse(strings.NewReader(test.args))
		if err != nil {
			log.Fatal(err)
		}
		words, images := countWordsAndImages(doc)
		if !reflect.DeepEqual(map[string]int{"words": words, "images": images}, test.want) {
			t.Errorf("%s", descr)
			t.Errorf("got-------------------")
			t.Errorf("words = %d\n", words)
			t.Errorf("images = %d\n", images)
			t.Errorf("expect---------------")
			for key, value := range test.want {
				t.Errorf("%s = %d", key, value)
			}
		}
	}
}
