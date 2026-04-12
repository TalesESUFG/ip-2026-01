package main

import f "fmt"

func main() {
	var (
		n int
	)
	f.Print("Informe o número de termos : ")
	f.Scan(&n)
	for i := 1; i <= n; i++ {
		p := i * i
		if i == n {
			f.Printf("%d.", p)
		} else {
			f.Printf("%d, ", p)
		}
	}
}
