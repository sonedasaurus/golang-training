package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args string
		want string
	}{
		{`<html><head><style src="/path/style"></style></head><body><a href="foo">Foo</a></body></html>`, "<html>\n  <head>\n    <style/ src=\"/path/style\">\n  </head>\n  <body>\n    <a href=\"foo\">\n      Foo\n    </a>\n  </body>\n</html>\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("outline(%q)", test.args)

		out = new(bytes.Buffer) // captured output
		doc, err := html.Parse(strings.NewReader(test.args))
		if err != nil {
			log.Fatal(err)
		}
		forEachNode(doc, startElement, endElement)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
