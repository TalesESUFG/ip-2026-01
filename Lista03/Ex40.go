package main

import f "fmt"

func main() {
	const per float64 = 0.6153846153846154
	var v float64
	for i := 6.0; i > 1; i -= 0.6 {
		f.Printf("Preço : %.1f ", i)
		v = i * 130
		f.Printf("Lucro esperado : %.2f ", v)
		v *= per
		f.Printf("Despesas esperadas : %.2f\n", v)
	}

}
