package main

import (
	f "fmt"
)

func main() {
	var (
		s   float64
		kw  float64
		gkw float64
		dc  float64
	)
	f.Print("Qual o valor do salário mínimo atual? ")
	f.Scan(&s)
	f.Print("Qual o gasto em kw do mês? ")
	f.Scan(&gkw)
	kw = (s * 0.7) / 100
	gkw *= kw
	dc = gkw * 0.9
	f.Printf("Preço pelo kW : R$ %.2f\nPreço da conta pelo kW gasto : R$ %.2f\nPreço da conta com desconto : R$ %.2f\n", kw, gkw, dc)
}
