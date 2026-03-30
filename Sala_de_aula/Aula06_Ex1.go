package main

import (
    f "fmt"
)

func main() {
    const NumNotas int = 3
    var (
        nota  [numNotas]float64
        soma  float64 = 0
        media float64
        m     float64
    )
    for i := 0; i < numNotas; i++ {
        f.Printf("Informe a nota%d : ", i)
        f.Scan(&nota[i])
        soma += nota[i]
    }
    m = 0
    for i, v := range nota {
        f.Printf("Nota %d = %f\n", i, v)
        m = m + v
    }
    media = m / float64(numNotas)
    f.Printf("Média das notas: %.2f\n", media)
    if media < 6 {
        f.Println("Você reprovou!")
    } else {
        f.Println("Você Passou!")
    }
}