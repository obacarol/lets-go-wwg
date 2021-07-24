package main

import "fmt"

/*
Faça um programa que:
	Declara uma variável;
	Atribui o valor de um número a ela;
	Informa se o número é positivo ou negativo.
*/

func main()  {
	var numero int64

	numero = -12

	if numero >= 0{
		fmt.Println("Número positivo (ou zero :p ).")
	}else{
		fmt.Println("Número negativo.")
	}
}
