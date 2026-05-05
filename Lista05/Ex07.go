package main

import f "fmt"

func main() {
	var (
		num [100]int
		v   []int
	)
	for i := len(num); i > 0; i-- {
		if i%2 != 0 {
			v = append(v, i)
		}
	}
	f.Print(v)
}
