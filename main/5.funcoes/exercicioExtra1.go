package main

import "fmt"

/*
Construa uma função que receba uma lista de números inteiros,
modifique essa lista dobrando os números ímpares
e divida por dois os pares,
e retorne a lista modificada e a soma de todos os elementos da lista.
*/

func dobraEDivide(lista []int64) []int64 {
	for i := range lista {
		if (lista[i] % 2) != 0 {
			lista[i] = lista[i] * 2
		} else {
			lista[i] = lista[i] / 2
		}
	}
	return lista
}

func main() {
	lista := []int64{1, 2, 3, 4, 5, 6}
	dobraEDivide(lista)
	fmt.Println(lista)
}
