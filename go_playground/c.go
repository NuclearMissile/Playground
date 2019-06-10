package main

import "fmt"

type T struct {
	s string
}

func change(t *T) {
	tt := T{"abc"}
	*t = tt
	t = &tt
}

func main() {
	var i *int
	i = new(int)
	*i = 1000
	fmt.Println(*i)
	/*	t := T{"123"}
		change(&t)
		fmt.Println(t)*/
}
