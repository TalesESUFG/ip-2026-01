package main

import (
	f "fmt"
)

func Proc(arr []int) (s1, s2, s3 int, pos1, pos2, pos3 int) {
	if len(arr) < 3 {
		panic("ERRO! Número inválido.")
	}
	if arr[0] < arr[1] {
		s1, s2 = arr[0], arr[1]
		pos1, pos2 = 0, 1
	} else {
		s1, s2 = arr[1], arr[0]
		pos1, pos2 = 1, 0
	}
	if arr[2] < s2 {
		if arr[2] < s1 {
			s3 = s2
			pos3 = pos2
			s2 = s1
			pos2 = pos1
			s1 = arr[2]
			pos1 = 2
		} else {
			s3 = s2
			pos3 = pos2
			s2 = arr[2]
			pos2 = 2
		}
	} else {
		s3 = arr[2]
		pos3 = 2
	}
	for i := 3; i < len(arr); i++ {
		val := arr[i]
		if val < s3 {
			s3 = val
			pos3 = i
			if val < s2 {
				s3 = s2
				pos3 = pos2
				s2 = val
				pos2 = i
				if val < s1 {
					s3 = s2
					pos3 = pos2
					s2 = s1
					pos2 = pos1
					s1 = val
					pos1 = i
				}
			}
		}
	}
	return
}
func main() {
	var (
		q    int
		l, j []int
	)
	f.Print("Quantos funcionários?(acima de 3) ")
	f.Scan(&q)
	for i := 0; i < q; i++ {
		var n int
		f.Printf("Qual o número do %dº empregado? ", i+1)
		f.Scan(&n)
		l = append(l, n)
		f.Printf("Qual o número de meses trabalhado pelo %dº empregado? ", i+1)
		f.Scan(&n)
		j = append(j, n)
	}
	s1, s2, s3, p1, p2, p3 := Proc(j)
	f.Print("O trabalhador : ", l[p1], " é o 1º mais recente, meses trabalhados : ", s1, "\nO trabalhador : ", l[p2], " é o 2º mais recente, meses trabalhados : ", s2, "\nO trabalhador : ", l[p3], " é o 3º mais recente, meses trabalhados : ", s3)
}
