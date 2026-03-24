programa
{
	inclua biblioteca Matematica --> m
	funcao inicio()
	{
		real n
		escreva("Qual a sua nota? ")
		leia(n)
		limpa()
		n = m.arredondar(n, 1)
		se (n >= 9)
		{
			escreva("Nota: ",n)
			escreva(" Conceito: A")
		}
		senao se (n >= 7.5 e n < 9)
		{
			escreva("Nota: ",n)
			escreva(" Conceito: B")
		}
		senao se (n >= 6.0 e n < 7.5)
		{
			escreva("Nota: ",n)
			escreva(" Conceito: C")
		}
		senao 
		{
			escreva("Nota: ",n)
			escreva(" Conceito: D")
		}
	}
}
