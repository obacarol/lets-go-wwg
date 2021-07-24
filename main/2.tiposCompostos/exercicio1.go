package main

import "fmt"

/*
Escreva um programa que contenha um array com 7 elementos do tipo string.
O valor de cada elemento deve ser o número do índice escrito por extenso.
Printe na tela o tipo do seu array e os valores de seus elementos.
*/

func main()  {
	numeros := [7]string{"zero", "um", "dois", "três", "quatro", "cinco", "seis"}

	fmt.Printf("Numeros é do tipo %T.\n", numeros)
	fmt.Println(numeros)
}
