package main

import f "fmt"

func main() {
	var (
		n1 float64
		n2 float64
	)
	f.Print("Digite dois números positivos e irei informar a subtração do primeiro pelo segundo : ")
	f.Scan(&n1,&n2)
	n1 -= n2
	f.Printf("O resultado da subtração é : %.2f\n", n1)
}
