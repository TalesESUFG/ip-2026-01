package main

import f "fmt"

func main() {
	var j, c [24]int
	f.Print("Olá! Bem vindo ao sistemas de vagas.\n")
	for i := 0; i != 1; {
		var q string
		f.Print("Você quer ver as vagas de janela ou corredor? (digite sair para sair do sistema)")
		f.Scan(&q)
		if q == "janela" || q == "janelas" || q == "Janelas" || q == "Janela" {
			f.Print("Vagas disponíveis : ")
			var (
				cont int
				re   string
			)
			for y := range 24 {
				if cont == 24 {
					f.Print("Não há vagas disponíveis.\n")
					break
				}
				if j[y] == 0 {
					f.Printf("J%d ", y+1)
				} else {
					cont += 1
				}
			}
			f.Print("\nIrá alugar alguma vaga? (s/n)")
			f.Scan(&re)
			if re == "n" || re == "N" {
				break
			}
			if re == "y" || re == "Y" {
				var resp int
				f.Print("Qual vaga (número)? ")
				f.Scan(&resp)
				if resp > 24 {
					f.Print("Erro! vaga inválida.\n")
					break
				}
				if j[resp-1] == 0 {
					j[resp-1] = 1
					f.Print("Vaga confirmada!")
				} else {
					f.Print("Erro! vaga já obtida.\n")
					break
				}
			}
		}
		if q == "corredor" || q == "corredores" || q == "Corredores" || q == "Corredor" {
			f.Print("Vagas disponíveis : ")
			var (
				cont int
				re   string
			)
			for y := range 24 {
				if cont == 24 {
					f.Print("Não há vagas disponíveis.\n")
					break
				}
				if c[y] == 0 {
					f.Printf("C%d ", y+1)
				} else {
					cont += 1
				}
			}
			f.Print("\nIrá alugar alguma vaga? (s/n)")
			f.Scan(&re)
			if re == "n" || re == "N" {
				break
			}
			if re == "y" || re == "Y" {
				var resp int
				f.Print("Qual vaga (número)? ")
				f.Scan(&resp)
				if resp > 24 {
					f.Print("Erro! vaga inválida.\n")
					break
				}
				if c[resp-1] == 0 {
					c[resp-1] = 1
					f.Print("Vaga confirmada!")
				} else {
					f.Print("Erro! vaga já obtida.\n")
					break
				}
			}
		}
		if q == "sair" || q == "Sair" {
			i = 1
			break
		}
	}
}
