package main

import (
	f "fmt"
)

func main() {
	var num int
	var l string
	f.Println("Escolha o número de sua viagem. (1 - Região Norte 2 - Região Nordeste 3 - Região Centro-Oeste 4 - Região Sul)")
	f.Scan(&num)
	f.Println("Inclui retorno?(s/n)")
	f.Scan(&l)
	if num == 1 || num == 2 || num == 3 || num == 4 && l == "s" || l == "n" {
		if num == 1 && l == "s" {
			f.Println("O valor da passagem é: R$ 900,00")
		} else if num == 1 && l == "n" {
			f.Println("O valor da passagem é: R$ 500,00")
		} else if num == 2 && l == "s" {
			f.Println("O valor da passagem é: R$ 650,00")
		} else if num == 2 && l == "n" {
			f.Println("O valor da passagem é: R$ 350,00")
		} else if num == 3 && l == "s" {
			f.Println("O valor da passagem é: R$ 600,00")
		} else if num == 3 && l == "n" {
			f.Println("O valor da passagem é: R$ 350,00")
		} else if num == 4 && l == "s" {
			f.Println("O valor da passagem é: R$ 550,00")
		} else if num == 4 && l == "n" {
			f.Println("O valor da passagem é: R$ 300,00")
		}
	} else {
		f.Println("ERRO!")
	}
}
