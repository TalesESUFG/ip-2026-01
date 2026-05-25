package main

import f "fmt"

func main() {
	var v [10]int
	for i := range v {
		f.Printf("Digite a %dº idade : ", i+1)
		f.Scan(&v[i])
	}
	m := make(map[int]int)
	for _, i := range v {
		m[i]++
	}
	mx := 0
	for _, f := range m {
		if f > mx {
			mx = f
		}
	}
	var mo []int
	for i, y := range m {
		if y == mx {
			mo = append(mo, i)
		}
	}
	f.Print("Moda: ", mo)
}
