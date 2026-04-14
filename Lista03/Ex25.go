package main

import f "fmt"

func main() {
        var r float64
        for i := 1.0; i != 16384; i *= 2 {
                n := 225.0
                m := i/n
                r += m
        }
        f.Print("O resultado é : ", r)
}