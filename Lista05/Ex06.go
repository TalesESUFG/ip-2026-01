package main

import f "fmt"

func main() {
	var (
		num [100]int
		v   []int
	)
	for i := len(num); i > 0; i-- {
		v = append(v, i)
	}
	f.Print(v)
}
