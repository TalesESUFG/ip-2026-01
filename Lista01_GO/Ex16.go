package main

import f "fmt"

func main() {
	var s float64
	f.Print("Digite seu salário : ")
	f.Scan(&s)
	if s <= 0 {
		f.Println("ERRO! Salário Inválido.")
	} else {
		if s < 300 {
			s *= 1.50
			f.Printf("Salário com reajuste : %.2f", s)
		} else {
			s *= 1.30
			f.Printf("Salário com reajuste : %.2f", s)
		}
	}
}
