package main

import "fmt"

/*
Represente uma lista de compras usando uma slice de string.
Declare-a usando literal de slice.
Posteriormente, adicione mais itens à lista.
Printe todas as etapas.
*/

func main()  {
	listaDeCompras := []string{"Banana", "Maça", "Uva", "Morango", "Kiwi"}
	fmt.Println(listaDeCompras)

	listaDeCompras = append(listaDeCompras, "Mamão", "Melancia")
	fmt.Println(listaDeCompras)
}
