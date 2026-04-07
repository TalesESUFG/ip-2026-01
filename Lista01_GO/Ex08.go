package main

import (
	f "fmt"
)

func main() {
	var (
		r float64
		a float64
	)
	f.Print("Informe o raio da lata : ")
	f.Scan(&r)
	f.Print("Informe a altura da lata : ")
	f.Scan(&a)
	l := custo(r, a)
	f.Printf("O valor do custo é : R$%.2f\n", l)
}
func custo(raio, altura float64) float64 {
	const pi = 3.14159
	ac := pi * (raio * raio)
	al := (2 * pi * raio * altura)
	va := ((2 * ac) + al) * 100
	return va
}
