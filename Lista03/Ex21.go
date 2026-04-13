package main

import f "fmt"

func main() {
	var b, n int
	f.Print("Digite o valor de b e o valor de n, respectivamente : ")
	f.Scan(&b, &n)
	if n <= 1 || b < 2 {
		f.Println("Erro! Números inválidos.")
	} else {
		r := 1
		for i := 0; i < n; i++ {
			r *= b
		}
		f.Print("O valor final é : %d", r)
	}
}
