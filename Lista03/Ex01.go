package main

import f "fmt"

func main() {
	var (
		b float64
		e int
		r float64 = 1
	)
	f.Print("Informe uma base real e um expoente inteiro, respectivamente : ")
	f.Scan(&b, &e)
	if e < 0 {
		b = 1 / b
		e = -e
	} // como estou arredondando por duas casas decimais, muito provavelmente não irá demonstrar a maioria das exponenciações negativas.
	for i := 0; i < e; i++ {
		r *= b
	}
	f.Printf("O resultado desta potência é : %.2f\n", r)
}
