package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		arg  string
		want string
	}{
		{"hoge fuga piyo $foo $bar $baz", "hoge fuga piyo foo_text bar_text baz"},
		{"$hoge fuga piyo foo $bar $baz", "hoge_text fuga piyo foo bar_text baz"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("(%q)", test.arg)
		god := expand(test.arg, testFunc)
		if !reflect.DeepEqual(god, test.want) {
			t.Errorf("%s", descr)
			t.Errorf("got-------------------")
			t.Errorf("id = %s\n", god)
			t.Errorf("expect---------------")
			t.Errorf("id = %s\n", test.want)
		}
	}
}

func testFunc(s string) string {
	if s == "foo" {
		return "foo_text"
	}
	if s == "bar" {
		return "bar_text"
	}
	if s == "hoge" {
		return "hoge_text"
	}
	return s
}
