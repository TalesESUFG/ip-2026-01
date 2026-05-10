package main

import f "fmt"

func fibo(x int) int {
	if x <= 1 {
		return x
	}
	return fibo(x-1) + fibo(x-2)
}
func main() {
	var sf []int
	for i := 0; i < 50; i++ {
		sf = append(sf, fibo(i))
	}
	f.Print(sf)
}
