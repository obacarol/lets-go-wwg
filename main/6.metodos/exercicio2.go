package main

import "fmt"

/*
Considere um cenário em que você precise regularmente trabalhar com slices de inteiros
e extrair a soma e média dos números dessa lista. Usando o que você aprendeu sobre métodos,
qual seria sua solução para esse problema em Go?
 */

type grupo []int

func main() {
	g := grupo{2, 4, 6, 8, 10}

	soma := g.somaElementos()
	fmt.Println("A soma dos elementos:", soma)

	media := g.mediaDosElementos()
	fmt.Println("A média é:", media)
}

func (g grupo) somaElementos() int {
	var soma int
	for _, numero := range g {
		soma += numero
	}
	return soma
}

func (g grupo) mediaDosElementos() int {
	media := g.somaElementos() / len(g)
	return media
}
