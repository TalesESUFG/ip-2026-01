package main

import f "fmt"

type Pessoa struct {
	nome   string
	peso   float64
	altura float64
	quant  int
}

func main() {
	var (
		
	)
	pe := Pessoa{}
	var q []Pessoa
	f.Print("Digite a quantidade de pessoas a serem analisadas: ")
	f.Scan(&pe.quant)
	for i := 1; i <= pe.quant; i++ {
		num := Pessoa{}
		f.Print("Digite seu nome: ")
		f.Scan(&num.nome)
		f.Print("Digite sua altura: ")
		f.Scan(&num.altura)
		num.peso = (72.7 * num.altura) - 58
		q = append(q, num)
	}
	for i := 0; i < len(q); i++ {
		f.Printf("Nome : %s\n", q[i].nome)
		f.Printf("Seu peso ideal : %2.f\n", q[i].peso)
	}
	f.Println("FIM!")
}
