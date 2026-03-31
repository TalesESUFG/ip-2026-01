package main

import (
    f "fmt"
)

func main() {
    const numint int = 5
    var (
        num  [numint] int
        soma int
    )
    for i := 0; i < numint; i++ {
        f.Printf("Informe o %d valor : ", i+1)
        f.Scan(&num[i])
        soma += num[i]
    }
    f.Printf("a soma dos %d números é : %d\n", numint, soma)
}