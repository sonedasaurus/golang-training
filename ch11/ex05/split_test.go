package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"a:b:c", ";", 1},
		{"ab,cd,ef,gh", ",", 4},
		{"test1-test2-test3-test4-test5", "-", 5},
		{"test", "", 4},
		{"", ":", 1},
		{"", "", 0},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if len(words) != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, len(words), test.want)
		}
	}

}
