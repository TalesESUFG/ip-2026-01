package main

import (
	f "fmt"
)

func main() {
	var a, b int
	f.Println("Qual o número a?")
	f.Scan(&a)
	f.Println("Qual o número b?")
	f.Scan(&b)
	if a%b == 0 {
		f.Printf("O número %d é divisível por %d!\n", a, b)
	} else {
		f.Printf("O número %d não é divisível por %d!\n", a, b)
	}
}
