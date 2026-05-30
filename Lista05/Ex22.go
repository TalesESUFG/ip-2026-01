package main

import f "fmt"

func main() {
	var c [3]int
	var s [3]float64
	for i := range c {
		f.Print("Digite o número de uma (ou mais) das dez contas :\n")
		f.Scan(&c[i])
		if c[i] > 999 {
			panic("Contas só podem haver 3 digitos.")
		}
	}
	checkup := make(map[int]int)
	for _, i := range c {
		checkup[i]++
	}
	for _, i := range checkup {
		valor := 1
		if i > valor {
			panic("Não pode haver contas repetidas.")
		}
	}
	for i := range c {
		f.Printf("Digite o saldo da conta %d : ", c[i])
		f.Scan(&s[i])
	}
	for i := 0; i != 4; {
		f.Print("--- 1. Efetuar depósito ---\n--- 2. Efetuar Saque ---\n--- 3. Consultar o ativo bancário ---\n--- 4. Finalizar o Programa ---\n")
		f.Print("Digite o número respectivo do comando desejado : ")
		f.Scan(&i)
		if i == 1 {
			var (
				z int
				k int
			)
			f.Print("Qual a conta bancária : ")
			f.Scan(&z)
			for y := range c {
				if c[y] == z {
					var saldo float64
					f.Print("Qual é o saldo adicionado : ")
					f.Scan(&saldo)
					s[y] += saldo
					f.Print("Operação finalizada.\n")
					break
				} else {
					k += 1
				}
				if k == len(c)-1 {
					f.Print("Nenhuma conta foi encontrada com este número.\n")
					break
				}
			}
		}
		if i == 2 {
			var (
				z int
				k int
			)
			f.Print("Qual a conta bancária : ")
			f.Scan(&z)
			for y := range c {
				if c[y] == z {
					var saque float64
					f.Print("Qual é o valor do saque : ")
					f.Scan(&saque)
					if saque > s[y] {
						f.Print("Saldo insuficiente.\n")
						break
					} else {
						s[y] -= saque
						f.Print("Operação finalizada.\n")
						break
					}
				} else {
					k += 1
				}
				if k == len(c)-1 {
					f.Print("Nenhuma conta foi encontrada com este número.\n")
					break
				}
			}
		}
		if i == 3 {
			var t float64
			for y := range s {
				t += s[y]
				if y == len(s)-1 {
					f.Print("Valor de todos os saldos : ", t, "\n")
					break
				}
			}
		}
		if i == 4 {
			break
		}
	}
}
