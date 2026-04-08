package main

import f "fmt"

func main() {
	var (
		n1 float64
		n2 float64
		n3 float64
	)
	f.Print("Informe três números e retornarei sua multiplicação : ")
	f.Scan(&n1,&n2,&n3)
	m := n1 * n2 * n3
	f.Printf("O resultado da multiplicação é : ", m)
}
