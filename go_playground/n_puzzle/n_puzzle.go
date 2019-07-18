package main

import (
	"errors"
	"fmt"
	"strings"
)

type Matrix [][]int

type NPuzzle struct {
	CurrMatrix   Matrix
	InitMatrix   Matrix
	Size         int
	Steps        []Step
	CurrX, CurrY int
	fmt.Stringer
	Evaluate func(m Matrix, s *Step) float64
}

type Step struct {
	X1, Y1, X2, Y2 int
	fmt.Stringer
}

func (s *Step) Reverse() *Step {
	ns := new(Step)
	ns.X1 = s.X2
	ns.Y1 = s.Y2
	ns.X2 = s.X1
	ns.Y2 = s.Y1
	return ns
}

func (np *NPuzzle) Reset() *NPuzzle {
	np.CurrMatrix = deepCopy(np.InitMatrix)
	np.Steps = make([]Step, 0, 16)
	return np
}

func (np *NPuzzle) Move(s *Step) bool {
	ok := np.CurrMatrix.move(s)
	if !ok {
		return false
	}
	np.CurrX = s.X2
	np.CurrY = s.Y2
	return true
}

func (m Matrix) Pos() (currX, currY int) {
	for indexi, i := range m {
		for indexj, j := range i {
			if j == 0 {
				return indexj, indexi
			}
		}
	}
	panic("")
}

func (np *NPuzzle) FmtSteps() string {
	sb := strings.Builder{}
	tempM := deepCopy(np.InitMatrix)
	for index, step := range np.Steps {
		sb.WriteString(fmt.Sprintf("Step %d: Move from (%d, %d) to (%d, %d).\n", index,
			step.X1, step.Y1, step.X2, step.Y2))
		if !tempM.move(&step) {
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
	//sb.WriteString("CurrMatrix:\n")
	//sb.WriteString(np.CurrMatrix.String() + "\n")
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
	for indexi, i := range np.CurrMatrix {
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
	if size < 2 || size > 15 {
		return nil, errors.New("input size should be in [2, 15]")
	}
	np := new(NPuzzle)
	np.InitMatrix = getRandMatrix(size)
	np.CurrMatrix = deepCopy(np.InitMatrix)
	np.CurrX, np.CurrY = np.CurrMatrix.Pos()

	//fmt.Printf("InitMatrix:    %p\n", &np.InitMatrix)
	//fmt.Printf("CurrMatrix: %p\n", &np.CurrMatrix)
	//fmt.Printf("Size:          %p\n", &np.Size)
	//fmt.Printf("Steps:         %p\n", &np.Steps)

	np.Size = size
	np.Steps = make([]Step, 0, 16)
	return np, nil
}
