package main

import (
	f "fmt"
	m "math"
)

func main() {
	const pi float64 = 3.145926535
	var (
		ar  float64
		vol float64
		op  int
		r   float64
		al  float64
	)
	f.Println("Escolha qual objeto que irá ser calculado(seu número): 1-cone reto; 2-cilindro; 3-esfera")
	f.Scan(&op)
	if op == 1 {
		f.Println("qual o raio e altura do objeto, respectivamente?")
		f.Scan(&r, &al)
		vol = (pi * r * r * al) / 3
		soma := (r * r) + (al * al)
		sqrt := m.Sqrt(soma)
		ar = pi * r * sqrt
		f.Printf("A área de superfície desse cone reto corresponde a : %.2f\n", ar)
		f.Printf("O volume desse cone reto corresponde a : %.2f\n", vol)
	} else if op == 2 {
		f.Println("qual o raio e altura do objeto, respectivamente?")
		f.Scan(&r, &al)
		vol = pi * r * r * al
		ar = 2 * pi * r * al
		f.Printf("A área de superfície desse cilindro corresponde a : %.2f\n", ar)
		f.Printf("O volume desse cilindro corresponde a : %.2f\n", vol)
	} else if op == 3 {
		f.Println("qual o raio do objeto?")
		f.Scan(&r)
		vol = (4 / 3) * pi * r * r * r
		ar = 4 * pi * r * r
		f.Printf("A área de superfície dessa esfera corresponde a : %.2f\n", ar)
		f.Printf("O volume dessa esfera corresponde a : %.2f\n", vol)
	} else {
		f.Println("Erro, objeto inválido!")
	}
}
