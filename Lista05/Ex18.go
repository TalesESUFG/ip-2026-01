package main

import f "fmt"

func main() {
	var l, k [10]int
	for i := range l {
		f.Printf("Digite o %dº número : ", i+1)
		f.Scan(&l[i])
	}
	for i := range l {
		for y := range l {
			if l[i] > l[y] {
				l[i], l[y] = l[y], l[i]
			}
		}
	}
	for i, y := len(l)-1, 0; i >= 0; i-- {
		k[y] = l[i]
		y++
	}
	f.Print(k)
}
