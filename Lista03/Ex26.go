package main

import f "fmt"

func fac3(x int) int {
	if x == 0 {
		return 1
	}
	return x * fac(x-1)
}
func main() {
	v := 100.0
	s := 0.0
	for i := 0; i < 20; i++ {
		k := float64(fac(i))
		s += (v / k)
		v -= 1
	}
	f.Print("O resultado é : ", s)
}
