package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	var tests = []struct {
		input IntSet
		want  int
	}{
		{intset(), 0},
		{intset(0), 1},
		{intset(0, 63), 2},
		{intset(0, 63, 64), 3},
		{intset(0, 63, 64, 127), 4},
		{intset(0, 63, 64, 127, 128), 5},
	}

	for _, test := range tests {
		fmt.Sprintf("Len()")
		if test.input.Len() != test.want {
			t.Errorf("intset.String() is %s", test.input.String())
			t.Errorf("got: %d", test.input.Len())
			t.Errorf("want: %d", test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		input  IntSet
		remove []int
		want   IntSet
	}{
		{intset(), []int{1}, intset()},
		{intset(2, 64), []int{2}, intset(64)},
		{intset(2, 64), []int{2, 64}, intset()},
	}

	for _, test := range tests {
		fmt.Sprintf("Remove()")
		for _, value := range test.remove {
			test.input.Remove(value)
		}
		if test.input.String() != test.want.String() {
			for _, value := range test.remove {
				t.Errorf("intset.Remove(%d)", value)
			}
			t.Errorf("got: String() %s", test.input.String())
			t.Errorf("want: String() %s", test.want.String())
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []IntSet{
		intset(),
		intset(1),
		intset(2, 64),
	}

	for _, test := range tests {
		test.Clear()
		fmt.Sprintf("Clear()")
		want := intset()
		if test.String() != want.String() {
			t.Errorf("got: String() %s", test.String())
			t.Errorf("want: String() %s", want.String())
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []IntSet{
		intset(),
		intset(1),
		intset(2, 64),
	}

	for _, test := range tests {
		fmt.Sprintf("Copy()")
		got := test.Copy()
		got.Add(893)
		if test.String() == got.String() {
			t.Errorf("copy faild")
			t.Errorf("original: String() %s", test.String())
			t.Errorf("copy: String() %s", got.String())
		}
	}
}

func intset(xs ...int) IntSet {
	var s IntSet
	for _, x := range xs {
		s.Add(x)
	}
	return s
}

func assertIntSetEquals(t *testing.T, actual, expected IntSet) {
	if len(actual.words) == 0 && len(expected.words) == 0 {
		return
	}

	if !reflect.DeepEqual(actual.String(), expected.String()) {
		t.Errorf("got %q\nwant %q", actual.String(), expected.String())
	}
}
