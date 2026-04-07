package main

import f "fmt"

func main() {
	var t int
	f.Print("Número de horas utilizando a charrete : ")
	f.Scan(&t)
	if t%3 == 0 {
		t = (t / 3) * 10
		f.Printf("O valor a pagar é : %.2f\n", float64(t))
	} else {
		c1 := (t % 3) * 5
		c2 := ((t - (t % 3)) / 3) * 10
		v := c1 + c2
		f.Printf("O valor a pagar é : %.2f\n", float64(v))
	}
}
