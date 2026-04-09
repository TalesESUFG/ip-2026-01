package main

import (
	f "fmt"
)

func main() {
	var c float64
	f.Print("Carlos, por gentliza informe seu salário : ")
	f.Scan(&c)
	j := c / 3
	i := 0
	for j < c {
		c *= 1.02
		j *= 1.05
		i++
	}
	f.Printf("Seu salário antes de João lhe passar : R$%.2f\nSalário de Jão quando lhe passar : R$%.2f\nTempo em meses para ele lhe passar : %d\n", c, j, i)
}
