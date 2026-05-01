package main

import f "fmt"

func main() {
	var l1, l2 []int
	for i := 0; i < 10; i++ {
		var n int
		f.Printf("Digite o %dº número da lista 1 : ", i+1)
		f.Scan(&n)
		l1 = append(l1, n)
	}
	for i := 0; i < 5; i++ {
		var n int
		f.Printf("Digite o %dº número da lista 2 : ", i+1)
		f.Scan(&n)
		l2 = append(l2, n)
	}
	f.Println(l1)
	f.Println(l2)
}
