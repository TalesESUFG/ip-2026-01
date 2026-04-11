package main

import f "fmt"

type notas struct {
	nota1    float64
	nota2    float64
	media    float64
	situação string
}

func main() {
	var n []notas
	var apr, rep, exa, al int
	var media float64
	f.Print("Quantos alunos serão analisados? ")
	f.Scan(&al)
	for i := 1; i <= al; i++ {
		aluno := notas{}
		f.Printf("\nDigite a 1º nota do aluno %d : ", i)
		f.Scan(&aluno.nota1)
		f.Printf("\nDigite a 2º nota do aluno %d : ", i)
		f.Scan(&aluno.nota2)
		aluno.media = (aluno.nota1 + aluno.nota2) / 2
		media += aluno.media
		if aluno.media <= 3 {
			aluno.situação = "Reprovado"
			rep += 1
		} else if aluno.media > 3 && aluno.media < 7 {
			aluno.situação = "Exame"
			exa += 1
		} else {
			aluno.situação = "Aprovado"
			apr += 1
		}
		n = append(n, aluno)
	}
	f.Println(" ")
	for i := 0; i < len(n); i++ {
		f.Printf("Média do %dº aluno : %.2f\n", i+1, n[i].media)
		f.Printf("Situação do %dº aluno : %s\n", i+1, n[i].situação)
		f.Println(" ")
	}
	media /= float64(len(n))
	f.Printf("Quantidade de alunos aprovados : %d\n", apr)
	f.Printf("Quantidade de alunos em exame : %d\n", exa)
	f.Printf("Quantidade de alunos reprovados : %d\n", rep)
	f.Printf("Média da classe : %.2f\n", media)
}
