package main

import (
	"fmt"
)

/*
Utilizando a palavra reservada var declare uma variável numérica do tipo int.
Atribua um valor a essa variável.
Printe na tela o seu valor.
Repita para 3 variáveis diferentes.
*/

func main() {
	var um, dois, tres int
	um, dois, tres = 1, 2, 3
	fmt.Printf("Um:%d \nDois:%d \nTrês:%d\n", um, dois, tres)
}
