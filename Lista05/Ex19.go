package main

import f "fmt"

func main() {
	var (
		l1 [10]int
		l2 [5]int
	)
	for i := range l1 {
		f.Print("Digite um número da lista 1 : ")
		f.Scan(&l1[i])
	}
	for i := range l2 {
		f.Print("Digite um número da lista 2 : ")
		f.Scan(&l2[i])
	}
	for i := range l1 {
		f.Print("Para o número ", l1[i], " : \n")
		for y := range l2 {
			if l1[i]%l2[y] == 0 {
				f.Print("Divisível por ", l2[y], " na posição ", y, "\n")
			} else if y == 4 {
				f.Print("Nenhum número é divisível por ", l1[i], ".\n")
			}
		}
	}
}
