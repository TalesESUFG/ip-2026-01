package main

import f "fmt"

func main() {
	var n int
	f.Print("Diga um número inteiro e verificarei se é divisível três e cinco : ")
	f.Scan(&n)
	if n%5 == 0 && n%3 == 0 {
		f.Println("O número é divisível.")
	} else if n <= 0 {
		f.Println("ERRO! Número menor ou igual a zero.")
	} else {
		f.Println("O número não é divisível.")
	}
}
