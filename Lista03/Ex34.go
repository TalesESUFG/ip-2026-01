package main

import f "fmt"

func Eucl(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func MMC(a, b int, inteiro ...int) int {
	r := a * b / Eucl(a, b)

	for i := 0; i < len(inteiro); i++ {
		r = MMC(r, inteiro[i])
	}

	return r
}
func main() {
	var a, b int
	f.Print("Digite o par de números e irei retornar o MMC : ")
	f.Scan(&a, &b)
	f.Printf("O MMC entre o número %d e %d é : %d", a, b, MMC(a, b))
}
