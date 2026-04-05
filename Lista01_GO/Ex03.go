package main

import (
	f "fmt"
)

func main() {
	var n1, n2, n3 int
	f.Print("Diga três números inteiros e irei fazer a concatenação(apenas números abaixo de 10) : ")
	f.Scan(&n1, &n2, &n3)
	if n1 < 10 && n2 < 10 && n3 < 10 {
		op1 := (float64(n1) + (0.1 * float64(n2)) + (0.01 * float64(n3))) * 100
		op2 := op1 * op1
		f.Printf("Composição dos números : %d\nQuadrado da composição : %d\n", int(op1), int(op2))
	} else {
		f.Println("Número inválido!")
	}
}
