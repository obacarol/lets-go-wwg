package main

import (
	"fmt"
)

/*
Crie um mapa chamado ano onde as chaves (keys) são os números dos meses do ano (ex: 1, 2) e
os valores (values) são os nomes dos meses.
Printe na tela os nomes do meses 1 e 12.
Printe na tela o tamanho do mapa ano.
*/

func main() {
	ano := make(map[int64]string)

	ano[1] = "Janeiro"
	ano[2] = "Fevereiro"
	ano[3] = "Março"
	ano[4] = "Abril"
	ano[5] = "Maio"
	ano[6] = "Junho"
	ano[7] = "Julho"
	ano[8] = "Agosto"
	ano[9] = "Setembro"
	ano[10] = "Outubro"
	ano[11] = "Novemrbro"
	ano[12] = "Dezembro"

	fmt.Println(ano[1], ano[12])
	fmt.Printf("O ano tem %v meses.", len(ano))
}
