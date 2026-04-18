package main

import f "fmt"

func main() {
	var n1, n2, q int
	f.Print("Digite o dois números inteiros (dividendo e divisor, respectivamente) : ")
	f.Scan(&n1, &n2)
	for n1 > n2 {
		n1 -= n2
		q += 1
	}
	f.Println("O resto é :", n1)
	f.Println("O quociente é :", q)
}
