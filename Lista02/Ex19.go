package main

import (
	f "fmt"
	m "math"
)

func main() {
	const pi float64 = 3.14
	var (
		ar  float64
		vol float64
		op  int
		r   float64
		al  float64
	)
	f.Println("Escolha qual objeto que irá ser calculado: 1-cone reto; 2-cilindro; 3-esfera")
	f.Scan(&op)
	if op == 1 {
		f.Println("qual o raio e altura do objeto, respectivamente?")
		f.Scan(&r, &al)
		vol = (pi * r * r * al) / 3
		soma := (r * r) + (al * al)
		sqrt := m.Sqrt(soma)
		ar = (pi * r * sqrt)
		f.Printf("A área de superfície desse cilindro reto corresponde a : %.2f\n", ar)
		f.Printf("O volume desse cilindro reto corresponde a : %.2f\n", vol)
	} else if op == 2 {

	} else if op == 3 {

	} else {
		f.Println("Erro, objeto inválido!")
	}
}
