package main

import "fmt"

/*
Faça um programa em que 3 variáveis recebem valores diferentes e
informa qual a variável com maior valor.
*/

func main()  {
	a, b, c := 1, 2, 3

	if a == b && b == c{
		fmt.Println("Os números são todos iguais.")
	} else if a >= b {
		if a >= c {
			fmt.Println("A é a maior.")
		} else {
			fmt.Println("C é a maior.")
		}
	} else if b >= c {
		fmt.Println("B é o maior.")
	} else {
		fmt.Println("C é o maior.")
	}
}
