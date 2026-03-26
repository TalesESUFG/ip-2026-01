package main

import (
	f "fmt"
	m "math"
)

func main() {
	var num float64
	f.Println("Qual o número de entrada?")
	f.Scan(&num)
	var sqr = m.Sqrt(num)
	f.Printf("A raiz quadrada de %.2f é: %.2f\n", num, sqr)
}
