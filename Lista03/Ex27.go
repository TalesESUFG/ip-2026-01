package main

import f "fmt"

func fac4(n int) int {
	if n > 1 {
		return n * fac4(n-1)
	}
	return 1
}

func main() {
	var x, v float64
	f.Print("Digite o número real que irá ser utilizado na função : ")
	f.Scan(&x)
	var q []float64
	r := 1.0
	for i := 0; i < 20; i++ {
		r *= x * x
		q = append(q, r)
	}
	v = 0
	for i := 0; i < 20; i++ {
		n := 2
		if i == 0 {
			v += 1
		} else if i%2 != 0 {
			v -= q[i] / float64(fac4(n))
			n += 2
		} else {
			v += q[i] / float64(fac4(n))
			n += 2
		}
	}
	f.Println("Cosseno de (x) :", v)
	f.Println("Diferença entre o número x e seu cosseno :", x-v)
}
