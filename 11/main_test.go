package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	var execPath = os.Args[0]
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{execPath}, execPath},
		{[]string{execPath, "one"}, execPath + " one"},
		{[]string{execPath, "one", "two"}, execPath + " one two"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%q)", test.args)

		out = new(bytes.Buffer) // captured output
		echo(test.args)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
