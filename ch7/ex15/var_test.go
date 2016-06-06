package eval

import (
	"reflect"
	"testing"
)

func TestVars(t *testing.T) {
	tests := []struct {
		expr string
		want []string
	}{
		{"sqrt(A / pi)", []string{"A", "pi"}},
		{"pow(x, 3) + pow(y, 3)", []string{"x", "y"}},
		{"5 / 9 * (F - 32)", []string{"F"}},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := expr.Vars()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s.Var(), got %q\n want %q\n", test.expr, got, test.want)
		}
	}
}
