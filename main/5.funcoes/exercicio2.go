package main

import "fmt"

/*
Continuando na solução do exercício anterior, agora faça com que a função
que cumprimenta seja capaz de receber o nome da pessoa que está sendo cumprimentada,
e printe o nome junto do cumprimento.
*/

func main() {
	nome := "Carol"
	cumprimenta2(nome)
}

func cumprimenta2(nome string) {
	fmt.Println("Bom dia", nome)
}
