package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Field struct {
	cells [][]bool
	w, h  int
}

type Life struct {
	a, b *Field
	w, h int
}

func NewField(w, h int) *Field {
	cells := make([][]bool, h)
	for i := range cells {
		cells[i] = make([]bool, w)
	}
	return &Field{
		cells: cells,
		w:     w,
		h:     h,
	}
}

func NewLife(w, h int) *Life {
	a := NewField(w, h)
	for i := 0; i < (w*h)/4; i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a,
		b: NewField(w, h),
		w: w,
		h: h,
	}
}

func (field *Field) Set(x, y int, b bool) {
	field.cells[y][x] = b
}

func (field *Field) IsAlive(x, y int) bool {
	x += field.w
	x %= field.w
	y += field.h
	y %= field.h
	return field.cells[y][x]
}

func (field *Field) Next(x, y int) bool {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && field.IsAlive(x+i, y+j) {
				alive++
			}
		}
	}
	return alive == 3 || (alive == 2 && field.IsAlive(x, y))
}

func (life *Life) Step() {
	for y := 0; y < life.h; y++ {
		for x := 0; x < life.w; x++ {
			life.b.Set(x, y, life.a.Next(x, y))
		}
	}
	life.a, life.b = life.b, life.a
}

func (life *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < life.h; y++ {
		for x := 0; x < life.w; x++ {
			b := byte(' ')
			if life.a.IsAlive(x, y) {
				b = byte('*')
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	life := NewLife(40, 15)
	for i := 0; i < 300; i++ {
		life.Step()
		fmt.Print("\x0c", life) // Clear screen and print field.
		time.Sleep(time.Second / 30)
	}
}
