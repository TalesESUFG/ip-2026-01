programa
{
	inclua biblioteca Matematica --> mat
	funcao inicio()
	{
		real a
		real b
		real c
		real d
		escreva("Qual os valores da matriz 2x2?(enter cada entrada) ")
		leia(a,b,c,d)
		limpa()
		real matr = (a*d)-(b*c)
		matr = mat.arredondar(matr, 2)
		escreva("O valor determinante é: ",matr)
	}
}
