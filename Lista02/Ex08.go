package main

import (
	f "fmt"
)

func main() {
	var num int
	f.Println("Diga um número inteiro e irei verficar se está entre 20 e 90.")
	f.Scan(&num)
	if num > 20 && num < 90 {
		f.Printf("O número %d está entre 20 e 90.\n", num)
	} else {
		f.Printf("O número %d não está entre 20 e 90.\n", num)
	}
}
