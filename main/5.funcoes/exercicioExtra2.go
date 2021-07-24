package main

import (
	"fmt"
	"strings"
)

/*
Construa uma função que receba uma palavra ou frase e uma letra,
e retorne o número de ocorrências da letra informada.
*/

func verifica(frase, letra string) int {
	qtd := strings.Count(frase, letra)
	return qtd
}

func main() {
	frase := "O dia está lindo."
	letra := "o"

	fmt.Println(verifica(frase, letra))
}
