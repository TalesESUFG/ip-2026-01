package main

import f "fmt"

func main() {
	f.Println("Qual é o tamanho dos lados?")
	var l1, l2, l3 float64
	f.Scan(&l1, &l2, &l3)
	//Algoritmo de valores
	if l1+l2 > l3 && l1+l3 > l2 && l2+l3 > l1 {
		if l1 > l2 && l1 < l3 || l2 < l1 && l2 > l3 || l3 < l1 && l3 > l2 {
			f.Println("Escaleno")
		} else if l1 != l2 && l1 == l3 || l3 != l1 && l3 == l2 {
			f.Println("Isóceles")
		} else {
			f.Println("Equilátero")
		}
	} else {
		f.Println("Triângulo inválido!")
	}
}
