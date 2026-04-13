package main

import (
	f "fmt"
	m "math"
)

func main() {
	var sena float64
	for i := 0.0; i <= 6.3; i += 0.1 {
		sena = i - ((m.Pow(i, 3)) / 6) + ((m.Pow(i, 5)) / 120) - ((m.Pow(i, 7)) / 5040)
		f.Print(sena, "\n")
	}
}
