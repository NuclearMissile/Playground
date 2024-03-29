package parser

import (
	"fmt"
	"math"
)

func (v Var) Eval(env Env) float64 {
	if value, ok := env[v]; ok {
		return value
	} else {
		panic(fmt.Sprintf("%s not found in env", v))
	}
}

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unknown unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unknown binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "ln":
		return math.Log(c.args[0].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "mod":
		return math.Mod(c.args[0].Eval(env), c.args[1].Eval(env))
	}
	panic(fmt.Sprintf("unknown function call: %q", c.fn))
}

func (p postUnary) Eval(env Env) float64 {
	switch p.op {
	case '!':
		return math.Gamma(p.x.Eval(env) + 1)
	}
	panic(fmt.Sprintf("unknown post-unary operator: %q", p.op))
}
