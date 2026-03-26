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
	if (a>b && a>c){
		f.Println("O Maior número é:",a)
	} else if (b>a && b>c) {
		f.Println("O Maior número é:",b)
	} else if (c>a && c>b) {
		f.Println("O maior número é:",c)
	}
	if (a<b && b>c) {
		f.Printf("O número intermediário é: %d\n",b)
		f.Printf("O menor número é: %d\n",c)
	}
}
