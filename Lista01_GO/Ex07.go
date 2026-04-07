package main

import (
	f "fmt"
)

func main() {
	var (
		fa  float64
		chu float64
	)
	f.Print("Diga a temperatura em Fahrenheit : ")
	f.Scan(&fa)
	f.Print("Digite o volume de chuva em polegadas : ")
	f.Scan(&chu)
	fa = ((5 * fa) - 160) / 9
	chu *= 25.4
	f.Printf("O valor da temperatura em Celcius : %.2f\nO volume de chuva em mm : %.2f\n", fa, chu)
}
