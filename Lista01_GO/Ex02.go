package main

import (
	f "fmt"
)

func main() {
	var (
		i   int = 0
		num int
		p   int
		pl  float64
		g   float64
		ar  float64
		c   float64
	)
	f.Println("Qual o número de jogos?")
	f.Scan(&num)
	for i < num {
		i++
		f.Printf("Diga qual é o público pagante total do jogo %dº : ", i)
		f.Scan(&p)
		f.Printf("A percentagem de público categoria popular do jogo %dº : ", i)
		f.Scan(&pl)
		pl /= 100
		f.Printf("A percentagem de público categoria geral do jogo %dº : ", i)
		f.Scan(&g)
		g /= 100
		f.Printf("A percentagem de público categoria arquibancada do jogo %dº : ", i)
		f.Scan(&ar)
		ar /= 100
		f.Printf("A percentagem de público categoria cadeiras do jogo %dº : ", i)
		f.Scan(&c)
		c /= 100
		pl *= float64(p)
		g *= float64(p)
		ar *= float64(p)
		c *= float64(p)
		l := valor(pl, g, ar, c)
		f.Printf("O valor total arrecadado do jogo %dº é : %.2f\n", i, l)
	}
}
func valor(pop, ge, arq, ca float64) float64 {
	pop *= 1
	ge *= 5
	arq *= 10
	ca *= 20
	s := pop + ge + arq + ca
	return s
}
