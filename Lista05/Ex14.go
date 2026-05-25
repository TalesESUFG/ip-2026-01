package main

import f "fmt"

func main() {
	var v1 = [10]int{0, 5, 4, 2, 1, 5, 3, 2, 5, 9}
	var v2 = [10]int{1, 5, 4, 2, 0, 5, 3, 2, 5, 9}
	var v3 []int
	for i := 0; i < 10; i++ {
		v3 = append(v3, v1[i])
		v3 = append(v3, v2[i])
	}
	f.Print(v3)
}
