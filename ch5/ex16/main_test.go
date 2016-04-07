package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		argsSeparate string
		argsStrings  []string
		want         string
	}{
		{" ", []string{"apple", "banana", "apricot"}, "apple banana apricot"},
		{" and ", []string{"apple", "banana", "apricot"}, "apple and banana and apricot"},
		{"	", []string{"apple", "banana", "apricot"}, "apple	banana	apricot"},
		{" ", []string{"apple"}, "apple"},
		{" ", []string{}, ""},
		{"test", []string{"", ""}, "test"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("getSource(%q)", test.argsStrings)
		god := stringJoin(test.argsSeparate, test.argsStrings...)
		if !reflect.DeepEqual(god, test.want) {
			t.Errorf("%s", descr)
			t.Errorf("got")
			t.Errorf("%s", god)
			t.Errorf("expect")
			t.Errorf("%s", test.want)
		}
	}
}
