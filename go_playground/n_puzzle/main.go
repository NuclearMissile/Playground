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
	np, err := Init(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("--------%d Puzzle--------\n", n)
	fmt.Printf("verbose: %t, n: %d, init pos: (%d, %d)\n", v, n, np.CurrX, np.CurrY)
	fmt.Println("------------------------")
	fmt.Println(np.String())
}
