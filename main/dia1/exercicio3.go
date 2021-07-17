package main

import (
"fmt"
)

func main() {
	var produto1, produto2, produto3 float64
	produto1 = 22.22
	produto2 = 33.33
	produto3 = 44.44

	var qtd1, qtd2, qtd3 float64
	qtd1 = 2.22
	qtd2 = 3.33
	qtd3 = 4.44

	valorCompra := (produto1 * qtd1) + (produto2 * qtd2) + (produto3 * qtd3)

	fmt.Printf("A compra é R$ %v", valorCompra)
}