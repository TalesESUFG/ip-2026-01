package main

import f "fmt"

func main() {
	var h int = 1
	for i := 1; i < 50; i++ {
		div := float64(h) / float64(i)
		f.Printf("%.4f\n", div)
		h += 2
	}
}
