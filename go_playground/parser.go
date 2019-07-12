package main

import (
	"fmt"
	"github.com/NuclearMissile/Playground/go_playground/parser"
	"math"
)

func main() {
	expr, err := parser.Parse("sin(pi/6)+ln(e)")
	fmt.Println(expr)
	if err != nil {
		fmt.Println(err)
	}
	env := parser.Env{
		"pi": math.Pi,
		"e":  math.E,
	}
	fmt.Println(expr.Eval(env))
}
