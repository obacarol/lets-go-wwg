package main

import "fmt"

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
