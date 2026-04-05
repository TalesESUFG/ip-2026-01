package main

import (
	f "fmt"
)

func main() {
	var n1, n2, n3 float64
	f.Println("Informe as três notas para calcularmos sua média:")
	f.Scan(&n1, &n2, &n3)
	l := media(n1, n2, n3)
	f.Printf("Sua média é : %.2f\n", l)
}
func media(nm1, nm2, nm3 float64) float64 {
	var m = (nm1 + nm2 + nm3) / 3
	return m
}
