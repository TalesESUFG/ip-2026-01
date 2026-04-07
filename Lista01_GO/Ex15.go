package main

import f "fmt"

func main() {
	var n int
	f.Print("Digite um número inteiro e irei informar todos os quadrados pares anteriores : ")
	f.Scan(&n)
	if n < 2 || n > 2000 {
		f.Println("ERRO! Número inválido.")
	} else {
		for i := 2; i <= n; i += 2 {
			e := i * i
			f.Printf("%d ^ 2 : %d\n", i, e)
		}
	}
}
