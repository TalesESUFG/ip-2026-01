package main

import (
	f "fmt"
	m "math"
)

func main() {
	var r, v float64
	for r >= 0 && r < 20 {
		r += 0.5
		v = float64((4 / 3)) * m.Pi * (m.Pow(r, 3))
		f.Printf("Para o raio de valor %.1f, o resultado é : %.4f\n", r, v)
	}
}
