package main

import (
	f "fmt"
)

func main() {
	var n int
	f.Print("Digite o número e irei transformá-lo em base 2 : ")
	f.Scan(&n)
	var q []int
	for n > 0 {
		r := n % 2
		n /= 2
		q = append(q, r)
	}
	f.Print("O resultado da conversão é : ")
	for i := len(q) - 1; i > -1; i-- {
		f.Print(q[i])
	}
}
