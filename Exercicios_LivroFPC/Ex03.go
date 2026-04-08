package main
import f "fmt"
func main() {
	var (
		n1 float64
		n2 float64
	)
	f.Print("Digite dois números reais e irei retornar sua divisão : ")
	f.Scan(&n1,&n2)
	if n2 <= 0 {
		f.Println("ERRO! Divisão por zero ou negativo.")
	} else {
		n1/=n2
		f.Printf("O resultado da divisão é : %.2f\n", n1)
	}
}
