package main

import "fmt"

func main()  {
	cores := [7]string{}
	arcoIris := [7]string{"vermelho", "laranja", "amarelo", "verde", "azul", "anil", "violeta"}

	cores = arcoIris

	fmt.Printf("O arco-iris tem essas cores: %s.\n", arcoIris)
	fmt.Printf("Essas são as cores do arco-iris: %s.\n", cores)
}
