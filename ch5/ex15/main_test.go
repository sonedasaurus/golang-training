package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		args []int
		want int
	}{
		{[]int{3}, 3},
		{[]int{1, 2, 3, 4}, 4},
		{[]int{-1, -2, 3, 4}, 4},
		{[]int{}, -9223372036854775808},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("getSource(%q)", test.args)
		god := max(test.args...)
		if !reflect.DeepEqual(god, test.want) {
			t.Errorf("%s", descr)
			t.Errorf("got")
			t.Errorf("%d", god)
			t.Errorf("expect")
			t.Errorf("%d", test.want)
		}
	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		args []int
		want int
	}{
		{[]int{3}, 3},
		{[]int{1, 2, 3, 4}, 1},
		{[]int{-1, -2, 3, 4}, -2},
		{[]int{}, 9223372036854775807},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("getSource(%q)", test.args)
		god := min(test.args...)
		if !reflect.DeepEqual(god, test.want) {
			t.Errorf("%s", descr)
			t.Errorf("got")
			t.Errorf("%d", god)
			t.Errorf("expect")
			t.Errorf("%d", test.want)
		}
	}
}
