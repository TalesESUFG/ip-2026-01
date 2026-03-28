package main

import (
	f "fmt"
)

func main() {
	var num int
	f.Println("qual a idade da pessoa para categoriza-lá?")
	f.Scan(&num)
	if num >= 0 && num <= 2 {
		f.Println("Categoria: Recém nascido.")
	} else if num >= 3 && num <= 11 {
		f.Println("Categoria: Criança.")
	} else if num >= 12 && num <= 19 {
		f.Println("Categoria: Adolescente.")
	} else if num >= 20 && num <= 55 {
		f.Println("Categoria: Adulto.")
	} else if num > 55 {
		f.Println("Categoria: Idoso.")
	} else {
		f.Println("ERRO!")
	}
}
