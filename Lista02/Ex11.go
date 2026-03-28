package main

import (
	f "fmt"
)

func main() {
	var x float64
	f.Println("qual o valor de x?")
	f.Scan(&x)
	x = (8 / (2 - x))
	f.Printf("O resultado da função é: %.2f\n", x)
}
