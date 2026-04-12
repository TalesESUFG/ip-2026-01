package main

import f "fmt"

func fibo(x int) int {
	if x <= 1 {
		return x
	}
	return fibo(x-1) + fibo(x-2)
}
func main() {
	var n int
	f.Print("Bem vindo! digite o termo de fibonacci : ")
	f.Scan(&n)
	for i := 0; i <= n; i++ {
		f.Printf("%d ", fibo(i))
	}
}
