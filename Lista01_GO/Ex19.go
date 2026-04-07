package main

import f "fmt"

func main() {
	var (
		n int
		t float64 = 0
	)
	f.Print("Informe um número inteiro positivo maior que 1 : ")
	f.Scan(&n)
	if n < 1 {
		f.Println("ERRO! Número inválido.")
	} else {
		for i := 1; i <= n; {
			o := (1 / float64(n))
			t += o
			n--
		}
		f.Printf("%.6f\n", t)
	}
}
