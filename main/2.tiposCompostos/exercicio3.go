package main

import "fmt"

/*
Declare duas slices de int e preencha 6 posições de cada uma.
Some as slices, formando uma 3ª slice.
Printe na tela o valor das três.
*/

func main()  {
	pares := []int64{2, 4, 6, 8, 10, 12}
	impares := []int64{1, 3, 5, 7, 9, 11}

	numeros := append(pares, impares...)

	fmt.Println("Pares:", pares)
	fmt.Println("Impares:", impares)
	fmt.Println("Tudo junto e misturado:", numeros)
}
