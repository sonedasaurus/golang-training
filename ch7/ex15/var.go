package eval

import "fmt"

func (v Var) Vars() []string {
	return []string{string(v)}
}

func (l literal) Vars() []string {
	return nil
}

func (u unary) Vars() []string {
	switch u.op {
	case '+':
		return u.x.Vars()
	case '-':
		return u.x.Vars()
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Vars() []string {
	switch b.op {
	case '+':
		return append(b.x.Vars(), b.y.Vars()...)
	case '-':
		return append(b.x.Vars(), b.y.Vars()...)
	case '*':
		return append(b.x.Vars(), b.y.Vars()...)
	case '/':
		return append(b.x.Vars(), b.y.Vars()...)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Vars() []string {
	switch c.fn {
	case "pow":
		return append(c.args[0].Vars(), c.args[1].Vars()...)
	case "sin":
		return c.args[0].Vars()
	case "sqrt":
		return c.args[0].Vars()
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
