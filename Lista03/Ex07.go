package main

import (
	f "fmt"
	s "slices"
)

func main() {
	var q []int
	f.Print("Olá! seja bem vindo.\n")
	for i := 1; i <= 100000; i++ {
		var e int
		f.Printf("Digite o %dº número inteiro: ", i)
		f.Scan(&e)
		if e == 30000 {
			break
		}
		q = append(q, e)
	}
	var sm, p, pm, im int
	mi := s.Min(q)
	ma := s.Max(q)
	for i := 0; i < len(q); i++ {
		sm += q[i]
		if q[i]%2 == 0 {
			p += q[i]
			pm += 1
		} else {
			im += 1
		}
	}
	mediap := p / pm
	var perim float64 = (float64(im) / float64(len(q))) * 100
	f.Println(" ")
	f.Printf("Quantidade de números : %d\n", len(q))
	f.Printf("Média dos números pares : %d\n", mediap)
	f.Printf("A soma de todos os números é : %d\n", sm)
	f.Printf("O maior número é : %d\n", ma)
	f.Printf("O menor número é : %d\n", mi)
	f.Printf("A percentagem de números ímpares : %.2f", perim)
}
