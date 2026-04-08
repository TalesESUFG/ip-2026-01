package main

import f "fmt"

func main() {
  var (
    sf float64
    v float64
  )
  f.Print("Digite o seu salário : ")
  f.Scan(&sf)
  f.Print("Digite suas vendas totais : ")
  f.Scan(&v)
  v *= 0.04
  sf += v
  f.Printf("Seu salário final é : R$%.2f\n", sf)
}
