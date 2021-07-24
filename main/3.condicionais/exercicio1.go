package main

import (
	"fmt"
	"time"
)

/*
Declare uma variável e atribua a ela o ano de nascimento de uma pessoa.
Declare uma variável e atribua a ela o ano atual.
Escreva um programa que verifique se no ano atual essa pessoa já está apta a votar e
que printe na tela uma mensagem informando caso positivo e caso negativo.
*/

func main()  {
	anoNascimento := 2008
	anoAtual := time.Now().Year()
	idade := anoAtual - anoNascimento

	if idade >= 16{
		fmt.Println("Pode votar!")
	} else {
		fmt.Printf("Ainda não pode votar. Espere mais %v anos.", 16-idade)
	}
}
