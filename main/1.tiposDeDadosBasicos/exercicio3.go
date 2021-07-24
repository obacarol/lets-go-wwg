package main

import (
"fmt"
)

/*
Declare variáveis para representar o preço dos itens do mercado conforme os valores da tabela.
Declare variáveis para representar a quantidade ou peso dos itens do mercado conforme os valores da tabela.
A banana e o abacate serão calculados pelo peso, a cerveja e o salgadinho serão calculados pela quantidade de unidades.
Printe na tela os valores totais de cada um dos itens e o preço total da compra.
*/

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
