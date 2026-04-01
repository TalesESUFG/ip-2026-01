package main

import (
	f "fmt"
)

func main() {
	var (
		id int
	)
	f.Println("Qual sua idade?")
	f.Scan(&id)
	if id < 16 {
		f.Println("Não Eleitor.")
	} else if id >= 16 && id < 18 || id >= 65 && id < 120 {
		f.Println("Eleitor facultativo.")
	} else if id >= 18 && id < 65 {
		f.Println("Eleitor obrigatório.")
	} else {
		f.Println("Idade inválida.")
	}
}
