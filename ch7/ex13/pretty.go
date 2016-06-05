package eval

import (
	"fmt"
	"strconv"
)

func (v Var) Pretty(_ bool) string {
	return string(v)
}

func (l literal) Pretty(_ bool) string {
	return strconv.FormatFloat(float64(l), 'g', -1, 64)
}

func (u unary) Pretty(_ bool) string {
	switch u.op {
	case '+':
		return " +" + u.x.Pretty(false)
	case '-':
		return " -" + u.x.Pretty(false)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Pretty(braceFlag bool) string {
	switch b.op {
	case '+':
		if braceFlag {
			return "(" + b.x.Pretty(false) + " + " + b.y.Pretty(false) + ")"
		} else {
			return b.x.Pretty(false) + " + " + b.y.Pretty(false)
		}
	case '-':
		if braceFlag {
			return "(" + b.x.Pretty(false) + " - " + b.y.Pretty(false) + ")"
		} else {
			return b.x.Pretty(false) + " - " + b.y.Pretty(false)
		}
	case '*':
		return b.x.Pretty(true) + " * " + b.y.Pretty(true)
	case '/':
		return b.x.Pretty(true) + " / " + b.y.Pretty(true)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Pretty(_ bool) string {
	switch c.fn {
	case "pow":
		return "pow(" + c.args[0].Pretty(false) + ", " + c.args[1].Pretty(false) + ")"
	case "sin":
		return "sin(" + c.args[0].Pretty(false) + ")"
	case "sqrt":
		return "sqrt(" + c.args[0].Pretty(false) + ")"
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
