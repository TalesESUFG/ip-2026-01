package main

import f "fmt"

func main() {
	var (
		ho int
		mi int
		se int
	)
	f.Print("Digite as horas, minutos e segundos, respectivamente, e irei transformar tudo em segundos : ")
	f.Scan(&ho, &mi, &se)
	if ho < 0 || mi < 0 || se < 0 {
		f.Println("ERRO! números informados são negativos.")
	} else {
		l := tempo(ho, mi, se)
		f.Printf("O tempo em segundos é : %d\n", l)
	}
}
func tempo(h, m, s int) int {
	h = (h * 60) * 60
	m *= 60
	soma := h + m + s
	return soma
}
