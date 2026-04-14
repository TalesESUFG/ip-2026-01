package main

import f "fmt"

func main() {
	var q []float64
	x := 225.0
	x2 := 29.0
	var r float64
	for i := 1.0; i <= 16384; i *= 2 {
		r = i / x
		x -= x2
		x2 -= 2
		q = append(q, r)
	}
	s := 0.0
	for z := 0; z < 14; z++ {
		s = q[0]
		if z%2 == 0 {
			s -= q[z+1]
		} else {
			s += q[z+1]
		}
	}
	f.Println("O valor final é : ", s)
}
