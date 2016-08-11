package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args       string
		wantCounts map[rune]int
		wantUTFLen [utf8.UTFMax + 1]int
	}{
		{"test", map[rune]int{'t': 2, 'e': 1, 's': 1}, [utf8.UTFMax + 1]int{0, 4, 0, 0}},
		{"test1 test2", map[rune]int{' ': 1, '2': 1, '1': 1, 't': 4, 'e': 2, 's': 2}, [utf8.UTFMax + 1]int{0, 11, 0, 0}},
		{"test1\ntest2", map[rune]int{'\n': 1, '2': 1, '1': 1, 't': 4, 'e': 2, 's': 2}, [utf8.UTFMax + 1]int{0, 11, 0, 0}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("visit(%q)", test.args)
		input := strings.NewReader(test.args)
		counts, len, _ := charcount(input)

		// Counts
		if !reflect.DeepEqual(counts, test.wantCounts) {
			t.Errorf("%s", descr)
			t.Errorf("got")
			for key, value := range counts {
				t.Errorf("rune\tcounts")
				t.Errorf("%q\t%d", key, value)
			}
			t.Errorf("expect")
			for key, value := range test.wantCounts {
				t.Errorf("rune\tcounts")
				t.Errorf("%q\t%d", key, value)
			}
		}

		// UTFLength
		if !reflect.DeepEqual(len, test.wantUTFLen) {
			t.Errorf("%s", descr)
			t.Errorf("got")
			for key, value := range len {
				t.Errorf("len\tcounts")
				t.Errorf("%d\t%d", key, value)
			}
			t.Errorf("expect")
			for key, value := range test.wantUTFLen {
				t.Errorf("len\tcounts")
				t.Errorf("%d\t%d", key, value)
			}
		}
	}
}
