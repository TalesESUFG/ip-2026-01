package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Escreva um número inteiro de 3 dígitos:")
	fmt.Scan(&num)
	if num >= 100 && num < 1000 {
		num = (num / 10) % 10
		fmt.Printf("O dígito na casa das dezenas é: %d\n", num)
	} else {
		fmt.Println("Número inválido.")
	}
}
