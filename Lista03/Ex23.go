package main

import f "fmt"

func main() {
	var n int
	f.Print("Digite o número de termos da expressão : ")
	f.Scan(&n)
	if n < 3 {
		f.Println("ERRO! número inválido.")
	} else {
		i2 := 1
		s := 0.0
		for i := n; i >= 1; i -= 3 {
			m := float64(i) / float64(i2)
			s += m
			i2 += 1
		}
		f.Print("O resultado é : ", s)
	}
}
