package main

import f "fmt"

func main() {
	var p, im, num []int
	for i := 0; i < 10; i++ {
		var y int
		f.Printf("Digite o %dº número : ", i+1)
		f.Scan(&y)
		num = append(num, y)
	}
	for i := 0; i < len(num); i++ {
		if num[i]%2 == 0 {
			p = append(p, num[i])
		} else {
			im = append(im, num[i])
		}
	}
	if len(p) > 0 {
		var smp int = 0
		for i := len(p) - 1; i >= 0; i-- {
			smp += p[i]
		}
		f.Print("Números pares digitados : ", p, "\n")
		f.Print("Soma dos números pares : ", smp, "\n")
	} else {
		f.Println("Não foi digitado nenhum número par.")
	}
	if len(im) > 0 {
		f.Print("Números ímpares digitados : ", im, "\n")
		f.Print("Quantidade de números ímpares : ", len(im), "\n")
	} else {
		f.Println("Não foi digitado nenhum número ímpar.")
	}
}
