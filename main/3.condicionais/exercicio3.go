package main

import (
	"fmt"
)

/*
Faça um programa que, dada a idade de uma pessoa, verifique se ela é menor de idade,
maior de idade ou idosa utilizando a instrução if.
*/

func main()  {
	idade := 12

	if idade < 18 && idade >= 0{
		fmt.Println("Menor de idade.")
	}else if idade >= 18 && idade <= 60{
		fmt.Println("Maior de idade.")
	}else if idade > 60{
		fmt.Println("Idoso.")
	}else{
		fmt.Println("Idade inválida.")
	}
}
