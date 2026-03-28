package main

import (
	f "fmt"
)

func main() {
	var d, m, a int
	f.Print("Digite o dia: ")
	f.Scan(&d)
	f.Print("Digite o mês: ")
	f.Scan(&m)
	f.Print("Digite o ano: ")
	f.Scan(&a)
	if d < 1 || d > 31 {
		f.Print("Dia inválido.")
	} else {
		if m == 1 {
			f.Print("A data é: ", d, " de Janeiro de ", a)
		} else if m == 2 {
			f.Print("A data é: ", d, " de Fevereiro de ", a)
		} else if m == 3 {
			f.Print("A data é: ", d, " de Março de ", a)
		} else if m == 4 {
			f.Print("A data é: ", d, " de Abril de ", a)
		} else if m == 5 {
			f.Print("A data é: ", d, " de Maio de ", a)
		} else if m == 6 {
			f.Print("A data é: ", d, " de Junho de ", a)
		} else if m == 7 {
			f.Print("A data é: ", d, " de Julho de ", a)
		} else if m == 8 {
			f.Print("A data é: ", d, " de Agosto de ", a)
		} else if m == 9 {
			f.Print("A data é: ", d, " de Setembro de ", a)
		} else if m == 10 {
			f.Print("A data é: ", d, " de Outubro de ", a)
		} else if m == 11 {
			f.Print("A data é: ", d, " de Novembro de ", a)
		} else if m == 12 {
			f.Print("A data é: ", d, " de Dezembro de ", a)
		} else {
			f.Print("Mês inválido.")
		}
	}
}
