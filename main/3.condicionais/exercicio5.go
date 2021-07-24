package main

import "fmt"

/*
Declare uma variável chamada hora e atribuir um valor int a ela.
Usando switch, descreva cases e printe na tela se a hora corresponde
ao período da manhã, tarde, noite ou madrugada.
*/

func main()  {
	hora := 17

	switch {
	case hora >= 6 && hora < 12:
		fmt.Println("Manhã.")
	case hora >= 12 && hora < 18:
		fmt.Println("Tarde.")
	case hora >= 18 && hora < 0:
		fmt.Println("Noite.")
	case hora >= 0 && hora < 6:
		fmt.Println("Madrugada.")
	default:
		fmt.Println("Hora inválida.")
	}
}
