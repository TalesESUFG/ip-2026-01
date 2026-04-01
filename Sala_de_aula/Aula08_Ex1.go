package main

import (
	f "fmt"
)

func main() {
	var (
		n1, n2, n3 int
	)
	f.Println("Escreva 3 números inteiros para categorizá-los do menor para o maior.")
	f.Scan(&n1, &n2, &n3)
	l := leitura(n1, n2, n3)
	f.Printf("O maior número é : %d\n", l)
}
func leitura(nu1, nu2, nu3 int) int {
	maior := nu1
	if nu2 > nu1 && nu2 > nu3 {
		maior = nu2
	} else if nu3 > nu1 && nu3 > nu2 {
		maior = nu3
	}
	return maior
}
