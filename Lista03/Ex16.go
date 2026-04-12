package main

import f "fmt"

func fettx(x int) int {
	if x <= 1 {
		return 1
	}
	if x%2 == 0 {
		return fettx(x-1) + fettx(x-2)

	} else {
		return fettx(x-1) - fettx(x-2)
	}
}
func fetty(y int) int {
	if y <= 1 {
		return 1
	}
	if y%2 == 0 {
		return fettx(y-1) + fettx(y-2)
	} else {
		return fettx(y-1) - fetty(y-2)
	}
}

func main() {
	var n1, n2 int
	f.Print("Digite os dois primeiros números inteiros : ")
	f.Scan(&n1, &n2)
	f.Print("Termos de x : ")
	for i := 0; i <= n1; i++ {
		f.Printf("%d ", fettx(i))
	}
	f.Println(" ")
	f.Print("Termos de y : ")
	for i := 0; i <= n2; i++ {
		f.Printf("%d ", fetty(i))
	}
}
