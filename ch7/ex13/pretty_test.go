package eval

import (
	"fmt"
	"testing"
)

func TestPretty(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"sqrt(A / pi)", "sqrt(A / pi)"},
		{"pow(x, 3) + pow(y, 3)", "pow(x, 3) + pow(y, 3)"},
		{"5 / 9 * (F - 32)", "5 / 9 * (F - 32)"},
		{"5 / 9 * F - 32", "5 / 9 * F - 32"},
		{"5 / 9 * F - 32 + 5", "5 / 9 * F - 32 + 5"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := expr.Pretty(false)
		if got != test.want {
			t.Errorf("%s.Pretty(false), got %q\n want %q\n", test.expr, got, test.want)
		}
	}
}
