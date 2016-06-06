package eval

import "fmt"

func (v Var) Var() []string {
	return []string{v}
}

func (l literal) Var() []string {
	return nil
}

func (u unary) Var() []string {
	switch u.op {
	case '+':
		return u.x.Var()
	case '-':
		return u.x.Var()
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) []string {
	switch b.op {
	case '+':
		return append(b.x.Var(), b.y.Var()...)
	case '-':
		return append(b.x.Var(), b.y.Var()...)
	case '*':
		return append(b.x.Var(), b.y.Var()...)
	case '/':
		return append(b.x.Var(), b.y.Var()...)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) []string {
	switch c.fn {
	case "pow":
		return append(c.args[0].Var(), c.args[0].Var()...)
	case "sin":
		return c.args[0].Var()
	case "sqrt":
		return c.args[0].Var()
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//!-Eval2
