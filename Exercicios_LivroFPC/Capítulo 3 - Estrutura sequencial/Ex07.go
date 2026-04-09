package main

import f "fmt"

func main() {
  var p float64
  f.Print("Digite seu peso : ")
  f.Scan(&p)
  p1 := p * 1.15
  f.Printf("Se você engordar 15 por cento, seu peso será : %.2f\n", p1)
  p2 := p * 0.80
  f.Printf("Se você emagrecer 20 por cento, seu peso será : %.2f\n", p2)
}
