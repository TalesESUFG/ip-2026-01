package main
import f "fmt"
func main() {
	var n,s[] float64
	for i := 0, i < 15; i++ {
		var r float64
		f.Printf("Digite o %dº número : ")
		f.Scan(&r)
		n = append(n, r)
	}
}