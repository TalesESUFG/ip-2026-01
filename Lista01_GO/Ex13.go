package main

import f "fmt"

func main() {
	var n float64
	f.Print("Digite sua nota e iremos lhe dizer o conceito que a mesma se encaixa : ")
	f.Scan(&n)
	if n < 0 || n > 10 {
		f.Println("ERRO! nota inválida.")
	} else {
		if n >= 9 {
			f.Println("Conceito A.")
		} else if n < 9 && n >= 7.5 {
			f.Println("Conceito B.")
		} else if n < 7.5 && n >= 6 {
			f.Println("Conceito C.")
		} else {
			f.Println("Conceito D.")
		}
	}
}
