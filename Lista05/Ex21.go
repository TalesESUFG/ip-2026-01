package main

import f "fmt"

func main() {
	var r [10]float64
	for i := range r {
		f.Printf("Digite o %dº número : ", i+1)
		f.Scan(&r[i])
	}
	for {
		var c int
		f.Print("código secreto : ")
		f.Scan(&c)
		switch c {
		case 0:
			return
		case 1:
			f.Print(r, "\n")
		case 2:
			var z []float64
			for i := len(r) - 1; i >= 0; i-- {
				z = append(z, r[i])
			}
			f.Print(z, "\n")
		}
	}
}
