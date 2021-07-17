package main

import (
	"fmt"
	"time"
)

func main() {
	var anoNascimento int64
	anoAtual := time.Now().Year()

	fmt.Printf("Qual ano você nasceu?\n")
	fmt.Scan(&anoNascimento)
	idade := int64(anoAtual) - anoNascimento
	fmt.Printf("Você tem %v anos.", idade)
}