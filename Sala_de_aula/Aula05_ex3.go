package main

import (
	f "fmt"
	m "math"
)

func main() {
	var al, i, nota, media, soma float64
	i = 1
	f.Print("Qual a quantidade de alunos?\n")
	f.Scan(&al)
	for i <= al {
		f.Print("Qual a nota do aluno ", i, "?\n")
		f.Scan(&nota)
		soma = nota + soma
		i++
	}
	media = (soma / al)
	m.Round(media)
	f.Printf("A média da sala é: %.2f \n", media)
}
