package main

import "fmt"

func main()  {
	listaDeCompras := []string{"Banana", "Maça", "Uva", "Morango", "Kiwi"}
	fmt.Println(listaDeCompras)

	listaDeCompras = append(listaDeCompras, "Mamão", "Melancia")
	fmt.Println(listaDeCompras)
}