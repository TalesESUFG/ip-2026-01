package main

import f "fmt"

func main() {
	var n, v, m int
	f.Print("Digite um número inteiro na base 8 e irei transformá-lo na base 10 : ")
	f.Scan(&n)
	var d []int
	var j []int
	for n > 0 {
		d = append(d, n%10)
		n /= 10
	}
	r := 1
	for i := len(d) - 1; i > -1; i-- {
		r *= 8
		v = r / 8
		j = append(j, v)
	}
	for i := len(d) - 1; i > -1; i-- {
		m += d[i] * j[i]
	}
	f.Print("O valor do número da base 8 para base 10 é : ", m)
}
