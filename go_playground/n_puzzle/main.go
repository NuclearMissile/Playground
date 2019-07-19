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
	flag.BoolVar(&v, "verbose", true, "verbose log flag")
	flag.IntVar(&n, "n", 3, "puzzle size")
}

func main() {
	flag.Parse()
	np, err := Init(n)
	if err != nil {
		fmt.Println(err)
		return
	}
	np.ObjFunc = MNorm1

	fmt.Printf("--------%d Puzzle--------\n", n)
	fmt.Printf("verbose: %t, n: %d, init pos of 0: (%d, %d)\n", v, n, np.CurrX, np.CurrY)
	fmt.Println("------------------------")
	np.Solve()
	fmt.Println(np)
}
