package main

import "fmt"

func main()  {
	numero := 18

	switch {
	case numero == 0:
		fmt.Println("Zero.")
	case numero % 2 == 0:
		if numero % 3 == 0{
			fmt.Println("Múltiplo de 2 e 3.")
		} else {
			fmt.Println("Múltiplo de 2.")
		}
	case numero % 3 == 0:
		fmt.Println("Múltiplo de 3.")
	default:
		fmt.Println("Nenhuma das opções.")
	}
}