package main

import f "fmt"

func main() {
	var mer float64
	i := 1
	m := 50
	var v int
	for m < 72 {
		v2 := m + 2
		i++
		v = m + v2
	}
	mer = float64(v) / float64(i)
	f.Printf("o resultado da soma dos valores pares é : %d\n", v)
	f.Printf("o resultado da média aritimética entre todos os valores pares do 50 ao 70 é : %.2f\n", mer)
}
