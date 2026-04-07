package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		al float64
		ar float64
	)
	f.Print("Digite a altura da pirâmide e aresta do hexágono em metros, respectivamente : ")
	f.Scan(&al, &ar)
	l := area(al, ar)
	f.Printf("O volume da pirâmide é : %.2f m³\n", l)
}
func area(al1, ar1 float64) float64 {
	sqr := m.Sqrt(3)
	ab := (3 * (ar1 * ar1) * sqr) / 2
	v := (al1 * ab) / 3
	return v
}
