package main

import (
	f "fmt"
	mr "math/rand/v2"
)

func main() {
	var v []int
	for i := 0; i < 20; i++ {
		n := mr.IntN(6) + 1
		v = append(v, n)
	}
	m := make(map[int]int)
	for _, i := range v {
		m[i]++
	}
	for c, v := range m {
		f.Print("Número ", c, " repetiu :  ", v, "\n")
	}
}
