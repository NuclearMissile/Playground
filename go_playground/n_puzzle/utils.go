package main

import (
	"bytes"
	"encoding/gob"
	"math"
	"math/rand"
	"time"
)

func getRandMatrix(size int) Matrix {
	l := getRandList(size * size)
	m := make([][]int, size, size)
	for i := range m {
		m[i] = make([]int, size)
		for j := range m[i] {
			m[i][j] = int(l[i*size+j])
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

func manhattan(s *Step) float64 {
	return math.Abs(float64(s.X1-s.X2)) + math.Abs(float64(s.Y1-s.Y2))
}

func euclid(s *Step) float64 {
	return math.Sqrt(math.Pow(float64(s.X1-s.X2), 2) + math.Pow(float64(s.Y1-s.Y2), 2))
}

func (m Matrix) move(step *Step) bool {
	if m[step.Y1][step.X1] != 0 || manhattan(step) != 1 {
		return false
	}
	m[step.Y1][step.X1], m[step.Y2][step.X2] =
		m[step.Y2][step.X2], m[step.Y1][step.X1]
	return true
}

func helper(m Matrix, s *Step, fn func(*Step) float64) []float64 {
	m.move(s)
	size := len(m)
	res := make([]float64, 0, size*size)
	for indexi, i := range m {
		for indexj, j := range i {
			var targetX, targetY int
			if j == 0 {
				targetX, targetY = size-1, size-1
			} else {
				targetY = int(j / size)
				targetX = j%size - 1
			}
			s := Step{}
			s.X1, s.Y1, s.X2, s.Y2 = indexj, indexi, targetX, targetY
			res = append(res, fn(&s))
		}
	}
	m.move(s.Reverse())
	return res
}

func norm1(l []float64) float64 {
	var res float64
	for _, i := range l {
		res += i
	}
	return float64(res)
}

func norm2(l []float64) float64 {
	var res float64
	for _, i := range l {
		res += i * i
	}
	return math.Sqrt(float64(res))
}

func MNorm1(m Matrix, s *Step) float64 {
	return norm1(helper(m, s, manhattan))
}

func MNorm2(m Matrix, s *Step) float64 {
	return norm2(helper(m, s, manhattan))
}

func ENorm1(m Matrix, s *Step) float64 {
	return norm1(helper(m, s, euclid))
}

func ENorm2(m Matrix, s *Step) float64 {
	return norm2(helper(m, s, euclid))
}
