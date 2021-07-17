package main

import (
"fmt"
)

func main() {
	//var produto1, produto2, produto3 float64
	banana := 1.25
	cerveja := 2.98
	abacate := 4.59
	salgadinho := 7.29

	qtdBanana := 2.17
	qtdCerveja := 6.0
	qtdAbacate := 5.65
	qtdSalgadinho := 3.0

	valorCompra := (banana * qtdBanana) + (cerveja * qtdCerveja) + (abacate * qtdAbacate) + (salgadinho * qtdSalgadinho)

	fmt.Printf("Banana = R$ %.4v\n", banana * qtdBanana)
	fmt.Printf("Banana = R$ %v\n", cerveja * qtdCerveja)
	fmt.Printf("Banana = R$ %.4v\n", abacate * qtdAbacate)
	fmt.Printf("Banana = R$ %v\n", salgadinho * qtdSalgadinho)
	fmt.Printf("Valor total da compra = R$ %.4v\n", valorCompra)
}