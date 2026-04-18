package main

import (
	f "fmt"
	m "math"
)

func main() {
	var g float64
	for i := 0; i < 64; i++ {
		if i == 0 {
			g = 1
		} else {
			g *= 2
		}
	}
	f.Println("O valor absoluto é :", m.Abs(g))
}
