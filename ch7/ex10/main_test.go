package main

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var tests = []struct {
		args StringSort
		want bool
	}{
		{StringSort{"apple", "banana", "apple"}, true},
		{StringSort{"apple", "banana", "lemon"}, false},
		{StringSort{"apple", "banana", "lemon", "banana", "apple"}, true},
		{StringSort{"apple"}, true},
		{StringSort{}, true},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("IsPalindrome(%q)", test.args)
		got := IsPalindrome(test.args)
		if got != test.want {
			t.Errorf("%s", descr)
			t.Errorf("got")
			t.Errorf("%d", got)
			t.Errorf("expect")
			t.Errorf("%d", test.want)
		}
	}
}
