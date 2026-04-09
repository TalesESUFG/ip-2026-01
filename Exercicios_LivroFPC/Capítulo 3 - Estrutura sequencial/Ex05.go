package main

import f "fmt"

func main() {
  var p float64
  f.Print("Digite o preço do produto : ")
  f.Scan(&p)
  p *= 0.9
  f.Printf("O novo preço do produto é : %.2f\n", p)
}
