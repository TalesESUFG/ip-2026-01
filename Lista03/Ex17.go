package main

import f "fmt"

func main() {
	var q []int
	for x := 1; x <= 10; x++ {
		q = append(q, x)
		s := 0
		s += 1
		s2 := 1
		for y := 1; y <= 10; y++ {
			s2 += 1
			q = append(q, y)
		}
		f.Print(q[s:s2], "\n")
	}
}
