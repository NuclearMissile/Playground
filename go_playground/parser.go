package main

import (
	"fmt"
	"math"

	"github.com/NuclearMissile/Playground/go_playground/parser"
)

func main() {
	expr, err := parser.Parse("ln(e)")
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
