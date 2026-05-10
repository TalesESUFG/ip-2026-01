package main

import (
	f "fmt"
	m "math"
)

func main() {
	var s []float64
	for i := 0; i < 15; i++ {
		var r, rs float64
		f.Printf("Digite o %dº número : ", i+1)
		f.Scan(&r)
		rs = m.Sqrt(r)
		s = append(s, rs)
	}
	f.Print(s)
}
