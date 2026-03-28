package main

import (
	f "fmt"
)

func main() {
	var num int
	f.Println("Diga um número inteiro e irei verficar se está entre 20 e 90.")
	f.Scan(&num)
	if num > 20 && num < 90 {
		f.Print("O número %d está entre 20 e 90.", num)
	}
}
