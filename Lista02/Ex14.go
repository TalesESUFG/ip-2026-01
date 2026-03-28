package main

import (
	f "fmt"
)

// testei usar vibecoding com o copilot aqui.
func main() {
	var p1 float64
	var mod string
	f.Print("Digite o valor do carro: ")
	f.Scan(&p1)
	f.Print("Irá querer uma modificação adicional?(s/n): ")
	f.Scan(&mod)
	if mod == "s" || mod == "n" {
		if mod == "n" {
			f.Printf("O valor do carro é: %.2f\n", p1)
		} else {
			f.Println("Escolha entre uma das modificações:\na - Ar condicionado (R$ 1750,00) \nb - Pintura metálica (R$ 800,00) \nc - Vidro elétrico (R$ 1200,00) \nd - Direção hidráulica (R$ 2000,00)\n")
			f.Scan(&mod)
			if mod == "a" {
				p1 += 1750
				f.Printf("O preço do carro é: %.2f", p1)
			} else if mod == "b" {
				p1 += 800
				f.Printf("O preço do carro é: %.2f", p1)
			} else if mod == "c" {
				p1 += 1200
				f.Printf("O preço do carro é: %.2f", p1)
			} else if mod == "d" {
				p1 += 2000
				f.Printf("O preço do carro é: %.2f", p1)
			} else {
				f.Print("Resposta inválida.")
			}
		}
	} else {
		f.Print("Resposta inválida.")
	}
}
