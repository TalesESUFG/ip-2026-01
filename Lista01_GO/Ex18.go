package main

import f "fmt"

func main() {
	var (
		vi int
		ra int
		ne int
	)
	f.Print("Digite o valor inicial, a razão e o número de elementos da PA, respectivamente : ")
	f.Scan(&vi, &ra, &ne)
	l := pa(vi, ra, ne)
	f.Printf("O resultado da PA é : %d\n", l)
}
func pa(v, r, n int) int {
	var s int = 0
	for i := 1; i <= n; i++ {
		p := v + ((i - 1) * r)
		s += p
	}
	return s
}
