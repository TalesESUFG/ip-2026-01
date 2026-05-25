package main

import f "fmt"

func main() {
	var l, l2 []int
	for i := 0; i < 30; i++ {
		var n int
		f.Print("Digite um número : ")
		f.Scan(&n)
		l = append(l, n)
	}
	for i := range l {
		var m int
		if l[i]%2 == 0 {
			m = l[i] * 2
			l2 = append(l2, m)
		} else {
			m = l[i] * 3
			l2 = append(l2, m)
		}
	}
	f.Print(l2)
}
