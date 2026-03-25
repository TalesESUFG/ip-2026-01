package main

import (
	f "fmt"
)

func main() {
	var num int
	f.Println("Qual o número para saber se é postivo, negativo ou nulo?")
	f.Scan(&num)
	if num == 0 {
		f.Println("O número é \"nulo\"!")
	} else if num >= 1 {
		f.Println("O número é positivo!")
	} else {
		f.Println("O númmero é negativo!")
	}
}
