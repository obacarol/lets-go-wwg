package main

import "fmt"

func main()  {
	numeros := [7]string{"zero", "um", "dois", "três", "quatro", "cinco", "seis"}

	fmt.Printf("Numeros é do tipo %T.\n", numeros)
	fmt.Println(numeros)
}