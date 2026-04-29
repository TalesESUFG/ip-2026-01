package main

import f "fmt"

func main() {
  var n[] int
  var m50[] int
  var l[] int
  for i := 1; i <= 10; i++{
    var m int
    f.Printf("Digite o %dº número : ",i)
    f.Scan(&m)
    n = append(n, m)
    if m > 50 {
      m50 = append(m50, m)
      l = append(l, i)
    }
  }
  f.Print("Todos os valores da lista : ",n,"\n")
  for i := 0; i < len(l); i++{
    f.Printf("O número %d está na posição : %d\n",m50[i],l[i])
  }
}
