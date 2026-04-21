package main

import (
	f "fmt"
)

func main() {
	var n int
	f.Print("Digite o número e irei transformá-lo em base 16 : ")
	f.Scan(&n)
	var q []int
	for n > 0 {
		r := n % 16
		n /= 16
		q = append(q, r)
	}
	f.Print("O resultado da conversão é : ")
	for i := len(q) - 1; i > -1; i-- {
		if q[i] == 10 {
			f.Print("A")
		} else if q[i] == 11 {
			f.Print("B")
		} else if q[i] == 12 {
			f.Print("C")
		} else if q[i] == 13 {
			f.Print("D")
		} else if q[i] == 14 {
			f.Print("E")
		} else if q[i] == 15 {
			f.Print("F")
		} else {
			f.Print(q[i])
		}
	}
}
