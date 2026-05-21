package main

import f "fmt"

func main() {
	var l1 []float64
	for i := range 15 {
		var l float64
		f.Printf("Digite a nota do %dº aluno : ", i+1)
		f.Scan(&l)
		if l > 10 {
			f.Print("ERRO! Número inválido!")
			break
		}
		l1 = append(l1, l)
	}
	m := make(map[float64]int)
	for _, i := range l1 {
		m[i] = m[i] + 1
	}
	for i := range l1 {
		for y := i + 1; y < len(l1); y++ {
			if l1[y] == l1[i] {
				l1[i] = 0
			}
		}
	}
	for i := range l1 {
		if m[l1[i]] > 1 {
			r := (float64(m[l1[i]]) / float64(len(l1))) * 100
			f.Printf("\nA nota %.2f repetiu %d vezes, e sua frequência relativa é : %.2f por cento", l1[i], m[l1[i]], r)
		}
	}

}
