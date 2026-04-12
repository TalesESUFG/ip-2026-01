package main

import f "fmt"

func main() {
	var n1, n2 int
	f.Printf("Digite dois números naturais e irei enctrar todos os primos entre eles : ")
	f.Scan(&n1, &n2)
	if n1 > n2 {
		for n2 < n1 {
			if n2%2 == 0 {
				n2 += 1
			} else {
				n2 += 2
			}
			if n2 > n1 {
				break
			}
			f.Printf("%d ", n2)
		}
	} else if n2 > n1 {
		for n1 < n2 {
			if n1%2 == 0 {
				n1 += 1
			} else {
				n1 += 2
			}
			if n1 > n2 {
				break
			}
			f.Printf("%d ", n1)
		}
	} else {
		f.Println("ERRO! Números inválidos.")
	}
}
