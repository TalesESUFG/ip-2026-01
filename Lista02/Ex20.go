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
	f.Println("Olá! Qual é o valor do produto e o método de pagamento?(Credito, dinheiro, cheque; ou parcela2x e parcela3x)")
	f.Scan(&valor, &con)
	if con == "Cheque" || con == "cheque" || con == "Dinheiro" || con == "dinheiro" {
		op = valor * 0.9
		f.Printf("O valor da compra será de : %.2f\n", op)
	} else if con == "Cartão" || con == "cartão" {
		op = valor * 0.95
		f.Printf("O valor da compra será de : %.2f\n", op)
	} else if con == "Parcela2x" || con == "parcela2x" {
		op = valor / 2
		f.Printf("O valor da compra será de 2x : %.2f\n", op)
	} else if con == "Parcela3x" || con == "parcela3x" {
		op = (valor * 1.1) / 3
		f.Printf("O valor da compra será de 3x : %.2f\n", op)
	} else if valor == 0 {
		f.Println("Erro! Valor inválido.")
	} else {
		f.Println("Erro! Método de pagamento inválido.")
	}
}
