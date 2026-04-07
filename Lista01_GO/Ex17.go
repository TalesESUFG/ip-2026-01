package main

import f "fmt"

func main() {
	var (
		n int
		q int
	)
	f.Print("Escreva um número par e, em seguida, a quantidade de números pares que queira obter após ele : ")
	f.Scan(&n, &q)
	if n%2 == 0 && n >= 0 {
		for i := 0; i <= q; i++ {
			f.Printf("%d ", n)
			n += 2
		}
	} else {
		f.Println("ERRO! Primeiro número não é par ou é menor que zero.")
	}
}
