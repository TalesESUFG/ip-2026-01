package main

import f "fmt"

func main() {
	var (
		num [10]int
	)
	for i := 0; i < len(num); i++ {
		f.Printf("Informe o %dº número : ", i+1)
		f.Scan(&num[i])
	}
	var mv, p int
	for i := 0; i < len(num); i++ {
		mv = num[i]
		for y := 0; y < len(num); y++ {
			if mv > num[y] {
				mv = num[y]
				p = y + 1
			}
		}
	}
	f.Print("O menor número é : ", mv, "\n")
	f.Print("A posição do número é : ", p, "\n")
}
