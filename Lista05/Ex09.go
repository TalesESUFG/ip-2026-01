package main

import f "fmt"

func main() {
	var (
		s         float64 = 0
		alt, altm []float64
	)
	for i := 0; i < 10; i++ {
		var a float64
		f.Printf("Digite a altura do %dº atleta : ", i+1)
		f.Scan(&a)
		s += a
		alt = append(alt, a)
	}
	s /= 10
	for i := 0; i < len(alt); i++ {
		if alt[i] > s {
			altm = append(altm, alt[i])
		}
	}
	f.Println("Média de altura : ", s)
	f.Println("Alturas acima da média : ", altm)
}
