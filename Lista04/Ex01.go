package main

import f "fmt"

func pot(b float64, e int) float64 {
	if e == 0 {
		return 1
	}
	return b * pot(b, e-1)
}
func main() {
	var (
		b float64
		e int
	)
	f.Print("Digite uma base e um expoente (positivo), respectivamente : ")
	f.Scan(&b, &e)
	r := pot(b, e)
	f.Println("O valor da potência é :", r)
}
