package main

import (
	f "fmt"
)

func main() {
	var c int
	var cs float64
	var t string
	f.Print("Qual sua conta,consumo em metros cúbicos e tipo?(inicial do tipo): ")
	f.Scan(&c, &cs, &t)
	if t == "r" || t == "R" {
		residencial := (1.05 * cs) + 5
		f.Printf("O valor da conta %d é: %.2f\n", c, residencial)
	} else if t == "c" || t == "C" && cs > 80 {
		cs = cs - 80
		comercial := (1.25 * cs) + 500
		f.Printf("O valor da conta %d é: %.2f\n", c, comercial)
	} else if t == "c" || t == "C" && cs <= 80 {
		comercial := 1.25 * cs
		f.Printf("O valor da conta %d é: %.2f\n", c, comercial)
	} else if t == "i" || t == "I" && cs > 100 {
		cs = cs - 100
		industrial := (1.04 * cs) + 800
		f.Printf("O valor da conta %d é: %.2f\n", c, industrial)
	} else if t == "i" || t == "I" && cs <= 100 {
		industrial := 1.04 * cs
		f.Printf("O valor da conta %d é: %.2f\n", c, industrial)
	} else {
		f.Println("Tipo de conta inválida.")
	}
}
