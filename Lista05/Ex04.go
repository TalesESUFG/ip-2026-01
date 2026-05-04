package main

import f "fmt"

func main() {
	var c, a []int
	q := 0
	for i := 10; i > 0; i-- {
		var n int
		q += 1
		f.Printf("Digite o %dº número : ", q)
		f.Scan(&n)
		a = append(a, n)
	}
	for i := len(a) - 1; i > 0; i-- {
		if i == 1 {
			break
		}
		if a[i] == a[i-1] {
			c = append(c, a[i])
		}
	}
	if len(c) > 0 {
		if len(c) > 1 {
			cr := 1
			for i := len(c) - 1; i > 0; i-- {
				if i == 1 {
					break
				}
				if c[i] == c[i-1] {
					cr += 1
				}
			}
			f.Print("Repetições : ", c, "\nQuantidade de vezes repetida : ", cr, "\n")
		} else {
			f.Print("Número repetido : ", c, "\n")
		}
	} else {
		f.Print("Não há nenhum número repetido na lista.")
	}
}
