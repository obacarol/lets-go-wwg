package main

import (
	"fmt"
)

/*
Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.
Passos:
	Criar um mapa com chave do tipo string e valor do tipo string (map[string]string) onde
		a chave seja o nome da cidade e o valor o nome do país.
	Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país.
	Esse mapa segundo mapa terá tipo map[string]int, sendo a chave o nome do país e o valor
		a quantidade de menções.
	Printe na tela os valores do último mapa.
*/

func main() {
	oceania := make(map[string]string)
	frequencia := make(map[string]int64)

	oceania["Sydney"] = "Australia"
	oceania["Melbourne"] = "Australia"
	oceania["Brisbane"] = "Australia"
	oceania["Camberra"] = "Australia"
	oceania["Adelaide"] = "Australia"
	oceania["Auckland"] = "Nova Zelandia"
	oceania["Wellington"] = "Nova Zelandia"
	oceania["Queenstown"] = "Nova Zelandia"
	oceania["Picton"] = "Nova Zelandia"
	oceania["Suva"] = "Fiji"
	oceania["Nadi"] = "Fiji"
	oceania["Ba"] = "Fiji"
	oceania["Nacualofa"] = "Tonga"
	oceania["Neiafu"] = "Tonga"

	frequencia["Australia"] = 0
	frequencia["Nova Zelandia"] = 0
	frequencia["Fiji"] = 0
	frequencia["Tonga"] = 0

	for k, _ := range frequencia {
		for _, v := range oceania {
			if k == v {
				frequencia[k]++
			}
		}
	}
	fmt.Println(frequencia)
}
