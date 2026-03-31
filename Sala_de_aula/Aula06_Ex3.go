package main

import (
    f "fmt"
)

func main() {
    var (
        num [10]int
    )
    for i := 0; i < len (num); i++ {
        f.Printf("Informe o %d valor : ", i+1)
        f.Scan(&num[i])
    }
    for i := 9; i >= 0; i-- {
        if i == 0 {
            f.Print(num[i], ". ")
        } else {
            f.Print(num[i], ", ")
        }
    }
}