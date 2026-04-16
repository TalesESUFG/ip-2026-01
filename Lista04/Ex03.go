package main

import f "fmt"

func reverso(s, e int, q []int) {
  if s > e{
    return
  }
  q[s], q[e] = q[e], q[s]
  reverso(s+1, e-1, q)
}
func main() {
	var e int
	var q[]int
	f.Print("informe o número de elementos : ")
	f.Scan(&e)
	for i:= 0; i < e;i++ {
		var n int
		f.Printf("Digite o %dº número : ", i+1)
		f.Scan(&n)
		q = append(q, n)
	}
	s := 0
	en := len(q)-1
	reverso(s, en , q)
	f.Println(q)
}
