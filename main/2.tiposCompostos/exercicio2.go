package main

import "fmt"

/*
Crie 2 arrays e atribua o valor do segundo ao primeiro.
*/

func main()  {
	cores := [7]string{}
	arcoIris := [7]string{"vermelho", "laranja", "amarelo", "verde", "azul", "anil", "violeta"}

	cores = arcoIris

	fmt.Printf("O arco-iris tem essas cores: %s.\n", arcoIris)
	fmt.Printf("Essas são as cores do arco-iris: %s.\n", cores)
}
