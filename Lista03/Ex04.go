package main

import (
	f "fmt"
	m "math"
)

func sqrt(x float64) float64 {
	if x < 2 {
		return x
	}
	y := x
	z := (y + (x / y)) / 2
	for m.Abs(y-z) >= 0.00001 {
		y = z
		z = (y + (x / y)) / 2
	}
	return z
}
func main() {
	var n int
	var q []float64
	var sn string
	f.Print("Digite quantos números inteiros irão ser avaliados : ")
	f.Scan(&n)
	for i := 1; i <= n; i++ {
		var num int
		f.Print("Digite o número : ")
		f.Scan(&num)
		q = append(q, float64(num))
	}
	for i := 0; i < len(q); i++ {
		l := sqrt(q[i])
		if m.Round(l) == m.Trunc(l) {
			sn = "Sim"
		} else {
			sn = "Não"
		}
		f.Printf("Numero %.f possui raiz quadrada perfeita : %s\n", q[i], sn)
	}
}
