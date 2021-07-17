package main

import (
	"fmt"
)

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