package main

import "fmt"

/*
Agora vamos mudar a abordagem: em vez da função printar o cumprimento,
ela vai gerar uma string com o cumprimento e retorná-la.
Quem a chamar é que vai decidir o que fazer com o cumprimento.
*/

func main() {
	nome := "Carol"
	hora := 5
	fmt.Println(cumprimenta4(nome, hora))
}

func cumprimenta4(nome string, hora int) string {
	saudacao := ""
	switch {
	case hora >= 0 && hora <= 5:
		saudacao = "Boa madrugada " + nome
	case hora >= 6 && hora <= 11:
		saudacao = "Bom dia " + nome
	case hora >= 12 && hora <= 18:
		saudacao = "Boa tarde " + nome
	case hora >= 19 && hora <= 23:
		saudacao = "Boa noite " + nome
	default:
		saudacao = "Olá "
	}
	return saudacao
}
