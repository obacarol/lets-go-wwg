package main

import "fmt"

func main()  {
	pares := []int64{2, 4, 6, 8, 10, 12}
	impares := []int64{1, 3, 5, 7, 9, 11}

	numeros := append(pares, impares...)

	fmt.Println("Pares:", pares)
	fmt.Println("Impares:", impares)
	fmt.Println("Tudo junto e misturado:", numeros)
}