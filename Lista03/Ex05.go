package main

import f "fmt"

type pessoas struct {
	idade  int
	peso   float64
	altura float64
}

func main() {
	var d []pessoas
	f.Println("Olá! Seja bem vindo.")
	for i, y := 0, 1; i != 1; y++ {
		p := pessoas{}
		f.Printf("Idade da %dº pessoa : ", y)
		f.Scan(&p.idade)
		f.Printf("Altura da %dº pessoa : ", y)
		f.Scan(&p.altura)
		f.Printf("Peso da %dº pessoa : ", y)
		f.Scan(&p.peso)
		d = append(d, p)
		var r string
		f.Print("Quer continuar a usar o programa? (s/n) ")
		f.Scan(&r)
		if r == "n" || r == "N" {
			i = 1
		}
	}
	var d1, d2, d3 int = 0, 0, 0
	var s float64 = 0
	for i := 0; i < len(d); i++ {
		if d[i].idade > 50 {
			d1++
		} else if d[i].idade <= 20 && d[i].idade >= 10 {
			s += d[i].altura
			d2++
		}
		if d[i].peso < 40 {
			d3++
		}
	}
	media := s / float64(d2)
	t := len(d)
	var per float64 = (float64(d3) / float64(t)) * 100
	f.Print(per)
	f.Printf("Quantidade de pessoas acima de 50 anos : %d\n", d1)
	f.Printf("Média de altura entre pessoas de 10 a 20 anos : %.2f\n", media)
	f.Printf("Percetagem de pessoas abaixo de 40 quilos comparado ao total : %.2f\n", per)
}
