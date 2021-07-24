package main

import "fmt"

/*
Ainda na função que cumprimenta, agora faça com que ela seja capaz de receber o valor da hora e
decidir o cumprimento com base no valor passado, de acordo com a tabela.
*/

func main() {
	nome := "Carol"
	hora := 5
	cumprimenta3(nome, hora)
}

func cumprimenta3(nome string, hora int) {
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
		saudacao = "Olá " + nome
	}
	fmt.Println(saudacao)
}

