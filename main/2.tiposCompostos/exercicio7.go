package main

import "fmt"

func main()  {
	type Pessoa struct {
		Nome string
		Idade int
		CorFavorita string
	}

	filhoCaçula := Pessoa{
		Nome : "Luis",
		Idade : 7,
		CorFavorita: "Preto",
	}
	filhoPrimogenito := Pessoa{
		Nome : "João",
		Idade : 13,
		CorFavorita: "Preto",
	}
	marido := Pessoa{
		Nome : "Jader",
		Idade : 34,
		CorFavorita: "Azul",
	}
	fmt.Printf("Os nomes das pessoas que moram comigo são %s, %s e %s.\n", filhoCaçula.Nome, filhoPrimogenito.Nome, marido.Nome)
	fmt.Printf("As idades deles são %v, %v e %v anos.\n", filhoCaçula.Idade, filhoPrimogenito.Idade, marido.Idade)
	fmt.Printf("As cores favoritas deles são %s, %s e %s.\n", filhoCaçula.CorFavorita, filhoPrimogenito.CorFavorita, marido.CorFavorita)
}