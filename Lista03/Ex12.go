package main

import f "fmt"

func fac2(x int) int {
	if x == 0 {
		return 1
	}
	return x * fac2(x-1)
}
func main() {
	var n float64
	f.Print("Digite um número N real e irei fazer o somatório: ")
	f.Scan(&n)
	s := 0.0
	for i := 1; i <= 20; i++ {
		r := n / float64(fac2(i))
		if i == 1 {
			s += n
		} else if i%2 == 0 {
			s += r
		} else {
			s -= r
		}

	}
	f.Println("O somatório é :", s)
}
