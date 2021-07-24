package main

import "fmt"

/*
Usando um literal, crie um mapa que tem como chaves nomes de cores e
como valores seu código hexadecimal.
Delete uma entrada do mapa.
Printe todas as etapas.
*/

func main()  {
	cores := map[string]string{
		"preto" : "#000000",
		"azul" : "#0000FF",
		"cyan" : "#00FFFF",
		"verde" : "#008000",
		"roxo" : "#A020F0",
	}
	fmt.Println(cores)

	delete(cores, "preto")
	fmt.Println(cores)
}
