package main

import (
	"fmt"
	"math"
)

func (np *NPuzzle) Solve() {
	var optF float64
	var optStep *Step
	optF = math.Inf(1)
	steps := np.CurrMatrix.candidateSteps(optStep)

	for true {
		for _, s := range steps {
			np.CurrMatrix.exchange(&s)
			currF := float64(len(np.Steps)) + np.ObjFunc(np.CurrMatrix)
			if currF < optF {
				optF = currF
				optStep = &s
			}
			np.CurrMatrix.exchange(&s)
		}

		if optStep != nil {
			np.Move(optStep)
			if np.CurrMatrix.isGoal() {
				return
			}
			if v {
				fmt.Printf("objfunc: %f, steps: %d\n", np.ObjFunc(np.CurrMatrix), len(np.Steps))
				fmt.Println(np.CurrMatrix)
			}
			steps = np.CurrMatrix.candidateSteps(optStep)
			optF = math.Inf(1)
		} else {
			panic("")
		}
	}
}

func (m Matrix) candidateSteps(lastStep *Step) []Step {
	temp := make([]Step, 0, 4)
	res := make([]Step, 0, 4)
	posX, posY := m.posOf0()
	up, down, left, right := Step{X1: posX, Y1: posY, X2: posX, Y2: posY},
		Step{X1: posX, Y1: posY, X2: posX, Y2: posY},
		Step{X1: posX, Y1: posY, X2: posX, Y2: posY},
		Step{X1: posX, Y1: posY, X2: posX, Y2: posY}
	up.Y2--
	down.Y2++
	left.X2--
	right.X2++
	temp = append(temp, up, down, left, right)
	for _, s := range temp {
		if s.validate(m, lastStep) {
			res = append(res, s)
		}
	}
	return res
}
