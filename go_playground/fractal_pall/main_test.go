package main

import (
	"image/color"
	"testing"
)

func benchmark(b *testing.B, f func(complex128) color.Color) {
	for i := 1; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkNewton(b *testing.B) {
	benchmark(b, newton)
}
