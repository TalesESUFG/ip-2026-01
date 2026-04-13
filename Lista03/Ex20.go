package main

import f "fmt"

func main() {
	for x := 0; x <= 10; x++ {
		for y := 0; y <= 10; y++ {
			if x > y {
				f.Printf("Linha : %d Coluna : %d\n", x, y)
			}
		}
	}
}
