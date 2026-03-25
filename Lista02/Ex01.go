package main

import (
	f "fmt"
)

func main() {
	var num int
	f.Println("Qual o número para saber se é par ou ímpar?")
	f.Scan(&num)
	if num%2 == 0 {
		f.Print("O número é par!")
	} else {
		f.Print("O número é impar!")
	}
}
