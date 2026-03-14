programa
{
	inclua biblioteca Matematica --> mat
	funcao inicio()
	{
		inteiro nc
		inteiro nv = 0
		escreva("Qual a quantidade de casos a ser analisada? ")
		leia(nc)
		limpa()
		enquanto(nv < nc)
		{
			nv = nv+1
			inteiro t
			real p
			real g
			real a
			real c
			escreva("Qual a quantidade de ingressos do jogo Nº ",nv," no total? ")
			leia(t)
			limpa()
			escreva("Qual a porcentagem de ingressos do jogo Nº ",nv," popular? ")
			leia(p)
			limpa()
			escreva("Qual a porcentagem de ingressos do jogo Nº ",nv," geral? ")
			leia(g)
			limpa()
			escreva("Qual a porcentagem de ingressos do jogo Nº ",nv," arquibancada? ")
			leia(a)
			limpa()
			escreva("Qual a porcentagem de ingressos do jogo Nº ",nv," cadeiras? ")
			leia(c)
			limpa()
			real calp = (p/100)*(t)*(1)
			real calg = (g/100)*(t)*(5)
			real cala = (a/100)*(t)*(10)
			real calc = (c/100)*(t)*(20)
			real soma = (calp+calg+cala+calc)
			soma = mat.arredondar(soma, 2)
			escreva("A Renda do jogo Nº ",nv," é = ",soma,"\n") //infelizmente o comando é apagado no loop, e não tem como salvar os dados de cada loop também (pelo menos que eu saiba).
		}
	}
}
