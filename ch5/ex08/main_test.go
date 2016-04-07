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
		doc string
		id  string
	}{
		{`<html><head><style id='test' src="/path/style"></style></head><body><a href="foo">Foo Bar Hoge</a></body></html>`, "test"},
		{`<html><head><style src="/path/style"></style></head><body><a href="foo">Foo Bar Hoge</a></body></html>`, "id is nothing"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("(%q)", test.doc)
		doc, err := html.Parse(strings.NewReader(test.doc))
		if err != nil {
			log.Fatal(err)
		}
		god := "id is nothing"
		node := ElementByID(doc, test.id)
		if node != nil {
			for _, a := range node.Attr {
				if a.Key == "id" {
					god = a.Val
				}
			}
		}
		if !reflect.DeepEqual(god, test.id) {
			t.Errorf("%s", descr)
			t.Errorf("got-------------------")
			t.Errorf("id = %s\n", god)
			t.Errorf("expect---------------")
			t.Errorf("id = %s\n", test.id)
		}
	}
}
