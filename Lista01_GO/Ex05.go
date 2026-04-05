package main

import (
	f "fmt"
)

func main() {
	var (
		c   int
		con float64
		t   string
	)
	f.Print("Olá! Qual o número de sua conta? ")
	f.Scan(&c)
	f.Print("Qual o consumo em m³? ")
	f.Scan(&con)
	f.Print("Escolha a opção de consumo de sua conta ('C' - Comercial; 'I' - Industrial; 'R' - Residencial) ")
	f.Scan(&t)
	r := preco(con, t)
	if r == 1 {
		f.Println("Erro! tipo inválido.")
	} else {
		f.Printf("Conta : %d\nValor da conta : R$ %.2f\n", c, r)
	}
}
func preco(cs float64, tipo string) float64 {
	if tipo == "C" || tipo == "c" {
		if cs <= 80 {
			co := 500.0
			return co
		} else {
			co := 500 + ((cs - 80) * 0.25)
			return co
		}
	} else if tipo == "I" || tipo == "i" {
		if cs <= 100 {
			in := 800.0
			return in
		} else {
			in := 800 + ((cs - 100) * 0.04)
			return in
		}
	} else if tipo == "R" || tipo == "r" {
		re := 5 + (cs * 0.05)
		return re
	} else {
		e := 1.0
		return e
	}
}
