package main

import f "fmt"

type bois struct {
	numero int
	peso   float64
}

func main() {
	var d []bois
	for i := 0; i < 90; i++ {
		b := bois{}
		f.Print("Identificação do boi : ")
		f.Scan(&b.numero)
		f.Print("Peso do boi (em kg) : ")
		f.Scan(&b.peso)
		d = append(d, b)
	}
	z := 0.0
	id := 0
	for i := 0; i < 90; i++ {
		if d[i].peso > z {
			z = d[i].peso
			id = d[i].numero
		}
	}
	f.Printf("O boi de número %dº é o mais pesado. (%.f kg)\n", id, z)
}
