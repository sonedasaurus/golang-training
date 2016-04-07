package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestHtmlTagCount(t *testing.T) {
	var tests = []struct {
		args string
		want map[string]int
	}{
		{`<html><head></head><body><a href="foo">Foo</a></body></html>`, map[string]int{"html": 1, "head": 1, "body": 1, "a": 1}},
		{`<html><head></head><body><ul><li><a href="/foo">Foo</a></li><li><a href="/bar">Bar</a></li></ul></body></html>`, map[string]int{"html": 1, "head": 1, "body": 1, "a": 2, "ul": 1, "li": 2}},
		{`<html><head></head><body><ul><li><a href="/foo">Foo</a></li><li><a href="/bar">Bar</a></li></ul><ul><li><a href="/hoge">Hoge</a></li><li><a href="/piyo">Piyo</a></li></ul></body></html>`, map[string]int{"html": 1, "head": 1, "body": 1, "a": 4, "ul": 2, "li": 4}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("htmlTagCount(%q)", test.args)
		doc, err := html.Parse(strings.NewReader(test.args))
		if err != nil {
			log.Fatal(err)
		}
		var counts map[string]int
		counts = map[string]int{}
		htmlTagCount(counts, doc)
		if !reflect.DeepEqual(counts, test.want) {
			t.Errorf("%s", descr)
			t.Errorf("got")
			for tagName, tagCount := range counts {
				t.Errorf("tagName = %s, tagCount = %d", tagName, tagCount)
			}
			t.Errorf("expect")
			for tagName, tagCount := range test.want {
				t.Errorf("tagName = %s, tagCount = %d", tagName, tagCount)
			}
		}
	}
}
