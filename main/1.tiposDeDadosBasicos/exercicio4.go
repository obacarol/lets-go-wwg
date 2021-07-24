package main

import (
	"fmt"
)

/*
Utilizando a palavra reservada var, declare uma variável do tipo string, e
atribua a ela o valor que corresponde a seu nome usando uma literal de string interpretada.
Utilizando a marmota, declare uma variável do tipo string,
e atribua a ela o valor que corresponde à sua cor favorita usando uma literal de string bruta.
Printe na tela os valores de todas as variáveis formando uma frase que faça sentido.
*/

func main() {
	nome := "Carol"
	cor := "preto"

	fmt.Printf("Meu nome é %s e gosto da cor %s.", nome, cor)
}
