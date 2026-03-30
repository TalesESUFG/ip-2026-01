package main

import (
	f "fmt"
)

func main() {
	var valor float64
	var dia, categoria string
	f.Println("Qual o valor do cd, o tipo (comum ou lançamento), e o dia que será alugado?")
	f.Scan(&valor, &categoria, &dia)
	if categoria == "comum" || categoria == "Comum" {
		if dia == "segunda" || dia == "Segunda" || dia == "segunda-feira" || dia == "Segunda-feira" {
			valor = valor * 0.60
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "terça" || dia == "Terça" || dia == "terça-feira" || dia == "Terça-feira" {
			valor = valor * 0.60
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "quarta" || dia == "Quarta" || dia == "quarta-feira" || dia == "Quarta-feira" {
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "quinta" || dia == "Quinta" || dia == "quinta-feira" || dia == "Quinta-feira" {
			valor = valor * 0.60
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "sexta" || dia == "Sexta" || dia == "sexta-feira" || dia == "Sexta-feira" {
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "sábado" || dia == "Sábado" || dia == "sábado-feira" || dia == "Sábado-feira" {
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "domingo" || dia == "Domingo" || dia == "domingo-feira" || dia == "Domingo-feira" {
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else {
			f.Println("Dia inválido.")
		}
	} else if categoria == "lançamento" || categoria == "Lançamento" {
		if dia == "segunda" || dia == "Segunda" || dia == "segunda-feira" || dia == "Segunda-feira" {
			valor = valor * 0.45
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "terça" || dia == "Terça" || dia == "terça-feira" || dia == "Terça-feira" {
			valor = valor * 0.45
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "quarta" || dia == "Quarta" || dia == "quarta-feira" || dia == "Quarta-feira" {
			valor = valor * 1.15
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "quinta" || dia == "Quinta" || dia == "quinta-feira" || dia == "Quinta-feira" {
			valor = valor * 0.45
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "sexta" || dia == "Sexta" || dia == "sexta-feira" || dia == "Sexta-feira" {
			valor = valor * 1.15
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "sábado" || dia == "Sábado" || dia == "sábado-feira" || dia == "Sábado-feira" {
			valor = valor * 1.15
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else if dia == "domingo" || dia == "Domingo" || dia == "domingo-feira" || dia == "Domingo-feira" {
			valor = valor * 1.15
			f.Printf("O preço final do DVD é: %.2f\n", valor)
		} else {
			f.Println("Dia inválido.")
		}
	} else {
		f.Println("Categoria inválida.")
	}
}
