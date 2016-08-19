package main

import "testing"

func TestHas(t *testing.T) {
	intset := intset(0, 63, 64, 127, 128)
	mapintset := mapintset(0, 63, 64, 127, 128)
	var tests = []struct {
		input int
	}{
		{0},
		{1},
		{63},
		{64},
		{65},
		{126},
		{127},
		{128},
	}

	for _, test := range tests {
		if intset.Has(test.input) != mapintset.Has(test.input) {
			t.Errorf("intset.Has(%d) is %t", test.input, intset.Has(test.input))
			t.Errorf("mapintset.Has(%d) is %t", test.input, mapintset.Has(test.input))
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		input []int
	}{
		{[]int{0}},
		{[]int{0, 64}},
		{[]int{0, 65}},
		{[]int{0, 64, 65}},
		{[]int{0, 64, 65, 127, 128}},
	}

	for _, test := range tests {
		intset := intset()
		mapintset := mapintset()
		for _, i := range test.input {
			intset.Add(i)
			mapintset.Add(i)
		}
		for _, i := range test.input {
			if intset.Has(i) == mapintset.Has(i) {
				t.Errorf("intset.Has(%d) is %t", i, intset.Has(i))
				t.Errorf("mapintset.Has(%d) is %t", i, mapintset.Has(i))
			}
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

func mapintset(xs ...int) *MapIntSet {
	s := NewMapIntSet()
	for _, x := range xs {
		s.Add(x)
	}
	return s
}
