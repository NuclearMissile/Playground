package main

import (
	"fmt"
	"github.com/NuclearMissile/Playground/go_playground/parser"
	"math"
)

func main() {
	expr, _ := parser.Parse("sin(pi/6)")
	env := parser.Env{}
	env["pi"] = math.Pi
	fmt.Println(expr.Eval(env))
}
