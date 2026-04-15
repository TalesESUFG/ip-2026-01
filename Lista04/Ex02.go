package main

import f "fmt"

func soma(q []float64, n int) float64 {
	if n == 0 {
		return 0
	}
	return q[n-1] + soma(q, n-1)
}
func main() {
	var q []float64
	f.Print("Digite números reais para uma array. Se quiser parar, digite 0.\n")
	for i := 1; i != 0; i++ {
		var n float64
		f.Printf("número %d : ", i)
		f.Scan(&n)
		if n == 0 {
			break
		}
		q = append(q, n)
	}
	r := soma(q, len(q))
	f.Println("A soma da lista é :", r)
}
