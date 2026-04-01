package main

import (
	f "fmt"
)

func main() {
	var (
		x, y, z float64
	)
	f.Println("Escreva 3 números reais para calcular sua média.")
	f.Scan(&x, &y, &z)
	l := media(x, y, z)
	f.Printf("A média dos números é: %.2f\n", l)
}
func media(num1, num2, num3 float64) float64 {
	var (
		m float64
	)
	m = (num1 + num2 + num3) / 3
	return m
}
