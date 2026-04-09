package main

import f "fmt"

func main() {
	var n int
	f.Print("Informe um número inteiro positivo : ")
	f.Scan(&n)
	if n <= 0 {
		f.Println("O número informado é negativo ou igual a zero.")
	} else {
		s := 0
		for i := 0; i*(i+1)*(i+2) < n; {
			i++
			if i*(i+1)*(i+2) == n {
				f.Printf("O número %d é triangular de %d * %d * %d.\n", n, i, i+1, i+2)
			}
			s = i
		}
		if s*(s+1)*(s+2) != n {
			f.Printf("O número %d não é triangular.\n", n)
		}
	}
}
