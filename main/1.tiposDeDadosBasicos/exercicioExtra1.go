package main

import (
	"fmt"
)

/*
Imagem de um código.
Descubra por que não compila
Erros de compilação nos ajudam a compreender o que precisamos consertar em nosso código.
O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida na tela.
*/

func main() {
	var quilometros int64
	quilometros = 150

	fmt.Println(quilometros)
}
