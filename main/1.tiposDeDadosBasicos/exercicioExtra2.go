package main

import (
	"fmt"
	"time"
)

/*
Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual
(para fins de arredondamento, considere que ela já fez aniversário no ano atual).
Como podemos pegar a informação do ano diretamente do console?
*/

func main() {
	var anoNascimento int64
	anoAtual := time.Now().Year()

	fmt.Printf("Qual ano você nasceu?\n")
	fmt.Scan(&anoNascimento)
	idade := int64(anoAtual) - anoNascimento
	fmt.Printf("Você tem %v anos.", idade)
}
