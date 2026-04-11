package main

import f "fmt"

func main() {
	var q [20]int
	for i := 0; i < len(q); i++ {
		f.Printf("Número : %d\n", i+1)
	}
}
