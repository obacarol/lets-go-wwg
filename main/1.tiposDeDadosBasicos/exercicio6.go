package main

import "fmt"

/*
Estabeleça uma comparação que utilize o operador lógico E (&&) duas vezes
Crie variáveis que representem condições para serem comparadas
Printe na tela o resultado da comparação (utilizando o operador lógico E) das variáveis
- antes de rodar seu programa, tente prever qual será o resultado.
*/

func main() {
	temPelo := true
	ehFofinha := true
	gostaDeColo := true

	ehAMimi := temPelo && ehFofinha && gostaDeColo

	fmt.Println(ehAMimi)
	fmt.Printf("%T", ehAMimi)
}
