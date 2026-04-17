package main

import f "fmt"

func dec(n int) {
	if n > 0 {
		dec(n / 2)
		f.Print(n % 2)
	}
}
func main() {
	var n int
	f.Print("informe o número inteiro : ")
	f.Scan(&n)
	if n == 0 {
		f.Print("0")
	} else {
		dec(n)
	}
	f.Print("\n")
}
