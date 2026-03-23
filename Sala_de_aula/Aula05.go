package main

import "fmt"

func main() {
	fmt.Println("Qual é o tamanho dos lados?")
	var l1, l2, l3 float64
	fmt.Scan(&l1, &l2, &l3)
	if l1+l2 > l3 && l1+l3 > l2 && l2+l3 > l1 {
		if l1 > l2 && l1 < l3 || l2 < l1 && l2 > l3 || l3 < l1 && l3 > l2 {
			fmt.Println("Escaleno")
		} else if l1 != l2 && l1 == l3 || l3 != l1 && l3 == l2 {
			fmt.Println("Isóceles")
		} else {
			fmt.Println("Equilátero")
		}
	} else {
		fmt.Println("Triângulo inválido!")
	}
}
