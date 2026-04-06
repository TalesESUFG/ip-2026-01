package main

import (
	f "fmt"
)

func main() {
	var (
		n int
		i int = 0
	)
	f.Print("Quantas conversões irá fazer? ")
	f.Scan(&n)
	for i < n {
		i++
		var temp float64
		f.Printf("Qual a temperatura em Fahrenheit da %dº operação? ", i)
		f.Scan(&temp)
		r := conversoes(temp)
		f.Printf("A temperatura de %.2f Fahrenheit em Celcius é : %.2f\n", temp, r)
	}
}
func conversoes(fhr float64) float64 {
	var conv = (5 * (fhr - 32)) / 9
	return conv
}
