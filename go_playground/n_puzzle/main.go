package main

import (
	"flag"
	"fmt"
)

var (
	v bool
	n int
)

func init() {
	flag.BoolVar(&v, "verbose", false, "verbose log flag")
	flag.IntVar(&n, "n", 4, "puzzle size")
}

func main() {
	flag.Parse()
	fmt.Println("--------N Puzzle--------")
	fmt.Printf("verbose: %t, n: %d\n", v, n)
	fmt.Println("------------------------")
	np, err := Init(n)
	if err != nil {
		fmt.Println(err)
		return
	}
	np.Solve()
	fmt.Println(np.String())
}
