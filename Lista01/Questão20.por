programa
{
	funcao inicio()
	{
		inteiro h
		inteiro m
		inteiro s
		escreva("Entre as horas, minutos e segundos respectivamente para ver o tempo total em segundos.(enter para cada entrada) ")
		leia(h,m,s)
		limpa()
		h = (h*60)*60
		m = m*60
		inteiro soma = h+m+s
		escreva("O tempo em segundo é: ",soma)
	}
}
