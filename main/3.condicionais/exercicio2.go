package main

import "fmt"

func main()  {
	var numero int64

	numero = -12

	if numero >= 0{
		fmt.Println("Número positivo (ou zero :p ).")
	}else{
		fmt.Println("Número negativo.")
	}
}
