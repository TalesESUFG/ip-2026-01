package main

import f "fmt"

func main() {
	var (
		n1 float64
		n2 float64
	)
	f.Print("Informe as duas notas e retornarei a média ponderada : ")
	f.Scan(&n1,&n2)
	if n1 < 0 || n2 < 0 {
		f.Println("ERRO! Notas negativas.")
	} else {
		mp := (n1 * 2) + (n2 * 3)/(2 + 3)
		f.Printf("Sua média ponderada é : %.2f\n", mp)
	}
}
