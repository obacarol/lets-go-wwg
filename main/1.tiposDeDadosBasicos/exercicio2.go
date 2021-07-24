package main

import (
	"fmt"
)

/*
Utilizando a marmota :=, declare duas variáveis: a e b.
Atribua os seguintes valores a elas, respectivamente: 230 e 27.
Crie variáveis para representar os resultados das operações matemáticas soma (a + b) e subtração (a - b).
Printe na tela os valores de todas as variáveis, um em cada linha.
*/

func main() {
	a, b := 230, 27
	soma := a + b
	subtracao := a - b
	fmt.Printf("A: %d \nB: %d \nSoma: %d \nSubtração: %d", a, b, soma, subtracao)
}
