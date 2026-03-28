package main

import (
	f "fmt"
)

func main() {
	var num float64
	f.Println("Diga o valor da compra e direi o valor de venda.")
	f.Scan(&num)
	if num < 10 && num > 0 {
		num = num * 0.7
		f.Println("O lucro de venda é:", num)
	} else if num >= 10 && num < 30 {
		num = num * 0.5
		f.Println("O lucro de venda é:", num)
	} else if num >= 30 && num < 50 {
		num = num * 0.4
		f.Println("O lucro de venda é:", num)
	} else if num >= 50 {
		num = num * 0.3
		f.Println("O lucro de venda é:", num)
	} else {
		f.Println("Número inválido!")
	}
}
