package main

import (
	f "fmt"
)

func main() {
	var num1, num2 int
	f.Println("Qual o número de entrada?")
	f.Scan(&num1, &num2)
	var soma = num1 + num2
	if soma > 20 {
		soma += 8
		f.Println("O valor é:", soma)
	} else {
		soma -= 5
		f.Println("O valor é:", soma)
	}
}
