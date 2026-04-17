package main

import (
	f "fmt"
	m "math"
)

func main() {
	var q []float64
	var s, r float64 = 0, 1
	for i := 0; i < 51; i++ {
		if i == 0 {
			r = 1
			q = append(q, r)
		}
		n := 1.0
		n += 2
		r *= n * n * n
		q = append(q, r)
	}
	for i := 0; i < 51; i++ {
		if i%2 == 0 {
			s += 1 / q[i]
		} else {
			s -= 1 / q[i]
		}
	}
	f.Println("O número aproximado de ∏ com 51 termos é :", m.Cbrt((s * 32)))
}
