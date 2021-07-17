package main

import (
	"fmt"
)

func main()  {
	idade := 12

	if idade < 18 && idade >= 0{
		fmt.Println("Menor de idade.")
	}else if idade >= 18 && idade <= 60{
		fmt.Println("Maior de idade.")
	}else if idade > 60{
		fmt.Println("Idoso.")
	}else{
		fmt.Println("Idade inválida.")
	}
}
