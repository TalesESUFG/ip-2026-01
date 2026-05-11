package main

import f "fmt"

func main() {
	var l1, l2 []float64
	for i := 0; i < 15; i++ {
		var l float64
		f.Printf("Digite a nota do %dº aluno : ", i+1)
		f.Scan(&l)
		if l > 10 {
			f.Print("ERRO! Número inválido!")
			break
		}
		l1 = append(l1, l)
	}
	for i := 0; i < len(l1); i++ {
		for y := len(l1) - 1; y > 0; y-- {
			if l1[i] == l1[y] {
				l2 = append(l2, l1[i])
			}
		}
	}
	f.Print(l2)
}
