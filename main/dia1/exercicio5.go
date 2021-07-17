package main

import "fmt"

func main() {
	signos := make([]string, 12)

	signos[0] = "Aquário"
	signos[1] = "Peixes"
	signos[2] = "Aries"
	signos[3] = "Touro"
	signos[4] = "Gêmeos"
	signos[5] = "Câncer"
	signos[6] = "Leão"
	signos[7] = "Virgem"
	signos[8] = "Libra"
	signos[9] = "Escorpião"
	signos[10] = "Sagitário"
	signos[11] = "Capricórnio"

	fmt.Println(signos)
	fmt.Println(signos[2:8])
}