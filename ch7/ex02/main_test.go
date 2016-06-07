package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var tests = []struct {
		args string
		want int64
	}{
		{"apple", 5},
		{"apple banana apricot", 20},
		{"apple\nbanana\napricot", 20},
		{"", 0},
		{" ", 1},
		{"\n", 1},
	}

	for _, test := range tests {
		var b bytes.Buffer
		w, count := CountingWriter(&b)
		descr := fmt.Sprintf("w.Write(%q)", test.args)
		w.Write([]byte(test.args))

		if *count != test.want {
			t.Errorf("%s", descr)
			t.Errorf("expected count is %d, but got is %d", test.want, *count)
		}
	}
}
