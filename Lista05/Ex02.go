package main

import f "fmt"

func main() {
	var l1, l2, par, impar, r1, r2 []int
	f.Print("Digite os 10 números da lista 1 : ")
	for i := 0; i < 10; i++ {
		var n int
		f.Scan(&n)
		l1 = append(l1, n)
	}
	f.Print("Digite o 5 números da lista 2 : ")
	for i := 0; i < 5; i++ {
		var n int
		f.Scan(&n)
		l2 = append(l2, n)
	}
	for i := 0; i < len(l1); i++ {
		if l1[i]%2 == 0 {
			par = append(par, l1[i])
		} else {
			impar = append(impar, l1[i])
		}
	}
	for i := 0; i < len(par); i++ {
		var m int
		for y := 0; y < len(l2); y++ {
			m += l2[y]
		}
		r1 = append(r1, m+par[i])
	}
	f.Println(r1)
	for i := 0; i < len(impar); i++ {
		var m int
		for y := 0; y < len(l2); y++ {
			m += l2[y]
		}
		r2 = append(r2, m+impar[i])
	}
	f.Println(r2)
}
