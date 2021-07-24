package main

import (
	"fmt"
)

/*
Faça um programa a partir das mesmas orientações do exercício acima,
mas utilizando switch em vez de if.
*/

func main()  {
	idade := 32

	switch {
	case idade < 18 && idade >= 0:
		fmt.Println("Menor de idade.")
	case idade >= 18 && idade <= 60:
		fmt.Println("Maior de idade.")
	case idade > 60:
		fmt.Println("Idoso.")
	default:
		fmt.Println("Idade inválida.")
	}
}
