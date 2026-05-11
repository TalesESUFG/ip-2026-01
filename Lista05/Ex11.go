package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		j []float64
		l []int
	)
	for i := 0; i < 100; i++ {
		l = append(l, i+1)
	}
	var y float64
	for i := float64(len(l)); i > 50; i-- {
		var s float64
		s = (m.Pow(y-i, 3))
		y += 1
		j = append(j, s)
	}
	f.Print(j)
}
