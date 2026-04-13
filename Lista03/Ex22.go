package main

import f "fmt"

func main() {
	var i2 float64 = 1.0
	var s float64 = 0
	for i := 38.0; i > 1; i-- {
		m := (i * (i - 1)) / i2
		s += m
		i2 += 1
	}
	f.Print("O valor é : ", s)
}
