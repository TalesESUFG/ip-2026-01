package main

import (
	f "fmt"
)

func main() {
	var (
		a float64
		b float64
		c float64
	)
	f.Print("Digite os três coeficientes (A, B, C), respectivamente : ")
	f.Scan(&a, &b, &c)
	l := bsk(a, b, c)
	f.Printf("O valor de Delta é : %.2f\n", l)
}
func bsk(a1, b1, c1 float64) float64 {
	bas := (b1 * b1) - (4 * a1 * c1)
	return bas
}
