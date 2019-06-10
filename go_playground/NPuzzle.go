package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type NPuzzle struct {
	Matrix [][]int
	Size   int
	Steps  []Step
}

type Step struct {
	X1, Y1, X2, Y2 int
}

func (np *NPuzzle) Reset() {
	np.Matrix = getRandMatrix(np.Size)
	np.Steps = nil
}

func (np *NPuzzle) IsSolved() bool {
	return np.Steps != nil
}

func (np *NPuzzle) Solve() {

}

func (np *NPuzzle) PrintSteps() {

}

func Init(size int) (*NPuzzle, error) {
	if size < 2 || size > 8 {
		return nil, errors.New(fmt.Sprintf("Size(%d) should be in the range [2, 8]", size))
	}
	np := NPuzzle{Size: size}
	np.Reset()
	return &np, nil
}

func (np *NPuzzle) print(s1, s2 string) {
	if len(s1) != 0 {
		fmt.Println(s1)
	}

	for i := range np.Matrix {
		fmt.Printf("%2d\n", np.Matrix[i])
	}

	if len(s2) != 0 {
		fmt.Println(s2)
	}
}

func getRandMatrix(size int) [][]int {
	l := getRandList(size * size)
	m := make([][]int, size, size)
	for i := range m {
		m[i] = make([]int, size)
		for j := range m[i] {
			m[i][j] = l[i*size+j]
		}
	}
	return m
}

func getRandList(l int) []int {
	r := make([]int, l)
	for i := range r {
		r[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	return r
}

func main() {
	var inputN int
	fmt.Println("--------N Puzzle--------")
	_, err := fmt.Scanln(&inputN)
	if err != nil {
		fmt.Println(err)
		return
	}
	np, err := Init(inputN)
	if err != nil {
		fmt.Println(err)
		return
	}
	np.print("Init Matrix:", "")
}
