programa
{
	funcao inicio()
	{
		real k
		escreva("Escreva um número natural e maior que 1: ")
		leia(k)
		limpa()
		se(k <= 1)
		{
			escreva("Número inválido!")
		}
		senao
		{
			inteiro i = 1
			real g = 0.0
			enquanto (i <= k)
			{
				real op = (1/k)
				k = k-1
				g = g+op
			}	
			escreva(g)
		}
	}
}
