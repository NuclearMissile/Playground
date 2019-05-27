package main

import "fmt"

type T struct {
	s string
}

func change(t *T) {
	tt := T{"abc"}
	t = &tt
}

func main() {
	t := T{"123"}
	change(&t)
	fmt.Println(t)
}
