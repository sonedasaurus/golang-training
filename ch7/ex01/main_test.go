package main

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var w WordCounter
	var tests = []struct {
		args string
		want WordCounter
	}{
		{"apple", 1},
		{"apple banana apricot", 3},
		{"apple\nbanana\napricot", 3},
		{"", 0},
		{" ", 0},
		{"\n", 0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("w.Write(%q)", test.args)
		w.Write([]byte(test.args))
		if w != test.want {
			t.Errorf("%s", descr)
			t.Errorf("got")
			t.Errorf("%d", w)
			t.Errorf("expect")
			t.Errorf("%d", test.want)
		}
	}
}

func TestLineCounter(t *testing.T) {
	var l LineCounter
	var tests = []struct {
		args string
		want LineCounter
	}{
		{"apple", 1},
		{"apple banana apricot", 1},
		{"apple\nbanana\napricot", 3},
		{"", 0},
		{"\n", 1},
		{"\n\n", 2},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("l.Write(%q)", test.args)
		l.Write([]byte(test.args))
		if l != test.want {
			t.Errorf("%s", descr)
			t.Errorf("got")
			t.Errorf("%d", l)
			t.Errorf("expect")
			t.Errorf("%d", test.want)
		}
	}
}
