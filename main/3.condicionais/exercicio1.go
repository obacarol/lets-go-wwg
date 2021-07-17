package main

import (
	"fmt"
	"time"
)

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
