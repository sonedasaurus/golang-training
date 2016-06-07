package main

import "testing"

func TestCountingWriter(t *testing.T) {
	var tests = []struct {
		expect tree
		want   string
	}{
		{tree{1, &tree{2, nil, &tree{3, nil, nil}}, &tree{4, nil, nil}}, "1 2 3 4"},
		{tree{1, &tree{2, nil, &tree{3, nil, nil}}, nil}, "1 2 3"},
		{tree{1, nil, nil}, "1"},
	}

	for _, test := range tests {
		got := test.expect.String()

		if got != test.want {
			t.Errorf("expected is %s, but got is %s", test.want, got)
		}
	}
}
