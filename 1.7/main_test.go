package main

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	var validTests = []struct {
		args []string
	}{
		{[]string{}},
		{[]string{"http://example.com"}},
		{[]string{"http://example.com", "http://google.com"}},
	}

	for _, test := range validTests {
		descr := fmt.Sprintf("fetch(%q)", test.args)

		err := fetch(test.args)
		if err != nil {
			t.Errorf("%s = %q, want nil", descr, err)
		}
	}

	var invalidTests = []struct {
		args []string
		want string
	}{
		{[]string{"http://invalid"}, "Get http://invalid: dial tcp: lookup invalid: no such host"},
		{[]string{"invalid://example.com"}, "Get invalid://example.com: unsupported protocol scheme \"invalid\""},
	}

	for _, test := range invalidTests {
		descr := fmt.Sprintf("fetch(%q)", test.args)

		err := fetch(test.args)
		if err.Error() != test.want {
			t.Errorf("%s = %q, want %q", descr, err, test.want)
		}
	}
}
