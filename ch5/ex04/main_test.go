package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestGetSource(t *testing.T) {
	var tests = []struct {
		args string
		want []string
	}{
		{`<html><head><style src="/path/style"></style></head><body><a href="foo">Foo</a></body></html>`, []string{"/path/style"}},
		{`<html><head><script src="/path/script"></script></head><body><a href="foo">Foo</a></body></html>`, []string{"/path/script"}},
		{`<html><head><img src="/path/img"></img></head><body><a href="foo">Foo</a></body></html>`, []string{"/path/img"}},
		{`<html><head><script src="/path/script"></script><style src="/path/style"></style><img src="/path/img"></img></head><body><a href="foo">Foo</a></body></html>`, []string{"/path/script", "/path/style", "/path/img"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("getSource(%q)", test.args)
		doc, err := html.Parse(strings.NewReader(test.args))
		if err != nil {
			log.Fatal(err)
		}
		links := getSource(nil, doc)
		if !reflect.DeepEqual(links, test.want) {
			t.Errorf("%s", descr)
			t.Errorf("got")
			for _, link := range links {
				t.Errorf("%s", link)
			}
			t.Errorf("expect")
			for _, link := range test.want {
				t.Errorf("%s", link)
			}
		}
	}
}
