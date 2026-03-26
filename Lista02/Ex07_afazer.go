package main

import (
	f "fmt"
)

func main() {
	var a, b, c int
	f.Println("Qual o número a?")
	f.Scan(&a)
	f.Println("Qual o número b?")
	f.Scan(&b)
	f.Println("Qual o número c?")
	f.Scan(&c)
	if a < b && b > c {
		f.Printf("O número intermediário é: %d\n", b)
		f.Printf("O menor número é: %d\n", c)
	} else if c < b && b > a {
		f.Printf("O número intermediário é: %d\n", b)
		f.Printf("O menor número é: %d\n", a)
	} else if c < a && a > b {
		f.Printf("O número intermediário é: %d\n", a)
		f.Printf("O menor número é: %d\n", b)
	} else if b < a && a > c {
		f.Printf("O número intermediário é: %d\n", a)
		f.Printf("O menor número é: %d\n", c)
	} else if b < c && c > a {
		f.Printf("O número intermediário é: %d\n", c)
		f.Printf("O menor número é: %d\n", a)
	} else if a < c && c > b {
		f.Printf("O número intermediário é: %d\n", a)
		f.Printf("O menor número é: %d\n", b)
	}
}
