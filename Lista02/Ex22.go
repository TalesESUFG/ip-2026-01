package main

import (
	f "fmt"
)

func main() {
	const (
		sm float64 = 788
		he float64 = 10
	)
	var (
		mat int
		hr  float64
	)
	f.Println("Olá! Qual sua matrícula?")
	f.Scan(&mat)
	f.Println("Quantas horas extras trabalhadas este mês?")
	f.Scan(&hr)
	slr := he * hr
	sb := (3 * sm) + slr
	if sb > 1500 && sb < 2000 {
		inss := sb * 0.12
		sl := sb - inss
		f.Printf("Matrícula : %d\nSalário bruto : R$ %.2f\nSalário Liquido : R$ %.2f\n", mat, sb, sl)
	} else if sb > 2000 {
		inss := sb * 0.20
		sl := sb - inss
		f.Printf("Matrícula : %d\nSalário bruto : R$ %.2f\nSalário Liquido : R$ %.2f\n", mat, sb, sl)
	} else {
		f.Printf("Matrícula : %d\nSalário bruto : R$ %.2f\nSalário Liquido : R$ %.2f\n", mat, sb)
	}
}
