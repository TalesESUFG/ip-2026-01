package main

import f "fmt"

func main() {
	var l [10]int
	for i := range l {
		f.Printf("Digite o %dº número : ", i+1)
		f.Scan(&l[i])
	}
	for i := range l {
		if l[i]%2 != 0 {
			f.Print("número : ", l[i], " posição : ", i+1, "\n")
		}
	}
}
