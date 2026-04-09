package main

import f "fmt"

func main() {
	var mer float64
	var i int = 0
	var v int = 0
	for m := 50; m <= 70; m += 2 {
		v += m
		i++
		f.Printf(" %d\n", v)
	}
	mer = float64(v) / float64(i)
	f.Printf("o resultado da soma dos valores pares é : %d\n", i)
	f.Printf("o resultado da média aritimética entre todos os valores pares do 50 ao 70 é : %.2f\n", mer)
}
