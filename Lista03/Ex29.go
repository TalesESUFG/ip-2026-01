package main

import f "fmt"

func main() {
	var n, s int
	f.Print("Informe o número inteiro N positivo : ")
	f.Scan(&n)
	if n > 0 {
		for i := 1; i < n; i++ {
			s += i
		}
		f.Printf("O somátorio de todos os números de 1 até %d é : %d\n", n, s)
	} else {
		f.Println("ERRO! O número necessita ser positivo.")
	}
}
