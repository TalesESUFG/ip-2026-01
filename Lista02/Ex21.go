package main

import (
	f "fmt"
)

func main() {
	var (
		id         int
		n1, n2, n3 float64
		m          float64
		cal        float64
	)
	f.Println("Escreva seu número de identificação.")
	f.Scan(&id)
	f.Printf("Olá aluno %d! Por favor, informe as três notas obtidas.\n", id)
	f.Scan(&n1, &n2, &n3)
	f.Println("Por último, por favor informe sua média das atividades.")
	f.Scan(&m)
	cal = ((n1) + (n2 * 2) + (n3 * 3) + (m)) / 7
	if cal >= 9 && cal <= 10 {
		f.Println("Conceito A! Você passou!")
	} else if cal < 9 && cal >= 7.5 {
		f.Println("Conceito B! Você passou!")
	} else if cal >= 6 && cal < 7.5 {
		f.Println("Conceito C! Você passou!")
	} else if cal < 6 && cal >= 4 {
		f.Println("Conceito D! Você reprovou!")
	} else if cal < 4 {
		f.Println("Conceito E! Você reprovou!")
	}
}
