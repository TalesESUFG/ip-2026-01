package main

import (
	f "fmt"
)

func main() {
	var num int
	f.Println("Qual o número que é ou não divisível por 5?")
	f.Scan(&num)
	if num%5 == 0 {
		f.Println("O número é divisível por 5!")
	} else {
		f.Println("O número não é divisível por 5!")
	}
}
