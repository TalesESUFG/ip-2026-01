programa
{
	inclua biblioteca Matematica --> mat
	funcao inicio()
	{
		real a
		real b
		real c
		escreva("Qual os valores da função de 2º grau em ordem? (enter a cada entrada)")
		leia(a,b,c)
		limpa()
		real bask = (b*b)-(4*a*c)
		bask = mat.arredondar(bask, 2)
		escreva("O valor de delta é: ",bask)
	}
}
