package main

import f "fmt"

func fac(x int) int {
	if x == 0 {
		return 1
	}
	return x * fac(x-1)
}
func main() {
	var n int
	f.Print("Digite um número inteiro para obter seu fatorial : ")
	f.Scan(&n)
	if n < 1 {
		f.Print("ERRO! número inválido.")
	} else {
		for i := 0; i <= n; i++ {
			if i == n {
				f.Printf("O fatorial é : %d", fac(i))
			}
		}
	}
}
