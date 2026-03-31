package main

import (
	f "fmt"
)

func main() {
	var (
		valor float64
		con   string
		op    float64
	)
	f.Println("Olá! Qual é o valor do produto e o método de pagamento?(A vista e credito, dinheiro, cheque; ou parcela 2x e parcela 3x)")
	f.Scan(&valor, &con)
	if con == "A vista" || con == "a vista" && con == "Cheque" || con == "cheque" || con == "Dinheiro" || con == "dinheiro" {
		op = valor * 0.9
		f.Printf("O valor da compra será de : %.2f\n", op)
	} else if con == "A vista" || con == "a vista" && con == "Cartão" || con == "cartão" {
		op = valor * 0.95
		f.Printf("O valor da compra será de : %.2f\n", op)
	} else if con == "Parcela 2x" || con == "parcela 2x" {
		op = valor / 2
		f.Printf("O valor da compra será de 2x : %.2f\n", op)
	} else if con == "Parcela 3x" || con == "parcela 3x" {
		op = (valor * 1.1) / 3
		f.Printf("O valor da compra será de 3x : %.2f\n", op)
	}
}
