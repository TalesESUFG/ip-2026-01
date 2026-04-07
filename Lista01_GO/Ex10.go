package main

import f "fmt"

func main() {
	var (
		n1 float64
		n2 float64
		n3 float64
		n4 float64
	)
	f.Print("Digite os quatro elementos da matriz bidimensional (A, B, C, D), respectivamente : ")
	f.Scan(&n1, &n2, &n3, &n4)
	l := det(n1, n2, n3, n4)
	f.Printf("O valor da matriz é : %.2f\n", l)
}
func det(x, y, z, g float64) float64 {
	vdet := (x * g) - (y * z)
	return vdet
}
