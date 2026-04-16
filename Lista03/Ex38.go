package main

import f "fmt"

func main() {
	var n, n2, s, s2 int = 0, 0, 0, 0
	cpf := []int{1, 1, 1, 4, 4, 4, 7, 7, 7}
	for i := 10; i > 1; i-- {
		m := cpf[n] * i
		n += 1
		s += m
	}
	dv := s % 11
	if dv < 2 {
		cpf = append(cpf, 0)
	} else {
		cpf = append(cpf, 11-dv)
	}
	for i := 11; i > 1; i-- {
		m := cpf[n] * i
		n2 += 1
		s2 += m
	}
	s2 += 9
	dv2 := s2 % 11
	if dv2 < 2 {
		cpf = append(cpf, 0)
	} else {
		cpf = append(cpf, 11-dv2)
	}
	f.Println(cpf)
}
