package main

import (
	f "fmt"
)

func main() {
	var x int
	f.Println("Escreva 1 números inteiro para calcular seu fatorial.")
	f.Scan(&x)
	if x > 20 {
		f.Println("Erro!")
	} else {
		l := fatorial(x)
		f.Printf("O fatorial é: %d\n", l)
	}
}
func fatorial(num1 int) int {
	var s int = 1
	for num1 > 1 {
		s *= num1
		num1--
	}
	return s
}
