package main

import (
	f "fmt"
	m "math"
)

func main() {
	var a, b, c float64
	f.Print("Digite o valor de a: ")
	f.Scan(&a)
	f.Print("Digite o valor de b: ")
	f.Scan(&b)
	f.Print("Digite o valor de c: ")
	f.Scan(&c)
	delta := m.Pow(b, 2) - 4*a*c
	if delta < 0 {
		f.Println("A equação é imaginária.")
	} else if delta == 0 {
		raiz := -b / (2 * a)
		f.Printf("A equação é raiz única: %.2f\n", raiz)
	} else {
		raiz1 := (-b + m.Sqrt(delta)) / (2 * a)
		raiz2 := (-b - m.Sqrt(delta)) / (2 * a)
		f.Printf("A equação tem raizes distintas: %.2f e %.2f\n", raiz1, raiz2)
	}
}
