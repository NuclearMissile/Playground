package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Matrix [][]uint8

type NPuzzle struct {
	CurrentMatrix Matrix
	InitMatrix    Matrix
	Size          int
	Steps         []Step
	fmt.Stringer
}

type Step struct {
	X1, Y1, X2, Y2 uint8
	fmt.Stringer
}

func (np *NPuzzle) Reset() *NPuzzle {
	np.CurrentMatrix = deepCopy(np.InitMatrix)
	np.Steps = make([]Step, 0, 16)
	return np
}

func (np *NPuzzle) FmtSteps() string {
	sb := strings.Builder{}
	tempM := deepCopy(np.InitMatrix)
	for index, step := range np.Steps {
		sb.WriteString(fmt.Sprintf("Step %d: Move %d from (%d, %d) to (%d, %d).\n", index,
			np.CurrentMatrix[step.Y1][step.X1], step.X1, step.Y1, step.X2, step.Y2))
		if !tempM.Move(step) {
			panic("")
		}
		sb.WriteString(tempM.String())
	}
	return sb.String()
}

func (np *NPuzzle) String() string {
	sb := strings.Builder{}
	sb.WriteString("InitMatrix:\n")
	sb.WriteString(np.InitMatrix.String() + "\n")
	sb.WriteString("CurrentMatrix:\n")
	sb.WriteString(np.CurrentMatrix.String() + "\n")
	if len(np.Steps) != 0 {
		sb.WriteString(fmt.Sprintf("Total steps: %d\n", len(np.Steps)))
		if v {
			sb.WriteString(np.FmtSteps())
		}
	}
	return sb.String()
}

func (m Matrix) String() string {
	sb := strings.Builder{}
	for _, i := range m {
		sb.WriteString(fmt.Sprintf("%2d\n", i))
	}
	return sb.String()
}

func (np *NPuzzle) IsSolved() bool {
	for indexi, i := range np.CurrentMatrix {
		offset := indexi * np.Size
		for indexj, j := range i {
			num := offset + indexj + 1
			if num == np.Size*np.Size {
				return true
			}
			if int(j) != num {
				return false
			}
		}
	}
	return true
}

func Init(size int) (*NPuzzle, error) {
	if size < 2 || size > 9 {
		return nil, errors.New("input size should be in [2, 10)")
	}
	np := new(NPuzzle)
	np.InitMatrix = getRandMatrix(size)
	np.CurrentMatrix = deepCopy(np.InitMatrix)

	//fmt.Printf("InitMatrix:    %p\n", &np.InitMatrix)
	//fmt.Printf("CurrentMatrix: %p\n", &np.CurrentMatrix)
	//fmt.Printf("Size:          %p\n", &np.Size)
	//fmt.Printf("Steps:         %p\n", &np.Steps)

	np.Size = size
	np.Steps = make([]Step, 0, 16)
	return np, nil
}

func getRandMatrix(size int) Matrix {
	l := getRandList(size * size)
	m := make([][]uint8, size, size)
	for i := range m {
		m[i] = make([]uint8, size)
		for j := range m[i] {
			m[i][j] = uint8(l[i*size+j])
		}
	}
	return m
}

func deepCopy(src Matrix) Matrix {
	buf := serialize(src)
	return deserialize(buf)
}

func serialize(input Matrix) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	_ = encoder.Encode(input)
	return buf.Bytes()
}

func deserialize(input []byte) Matrix {
	var res Matrix
	decoder := gob.NewDecoder(bytes.NewReader(input))
	_ = decoder.Decode(&res)
	return res
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
