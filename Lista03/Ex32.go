package main

import f "fmt"

func main() {
	var n1, n2, x int
	f.Print("Digite os dois números inteiros : ")
	f.Scan(&n1, &n2)
	for i := 0; i < n2; i++ {
		x += n1
	}
	f.Println(x)
}
