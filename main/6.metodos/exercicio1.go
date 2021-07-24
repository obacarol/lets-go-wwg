package main

import (
	"fmt"
	"math"
)

/*
Construa dois métodos: um que retorna a área e outro que retorna
o perímetro de uma estrutura que representa um círculo.
Printe a área e o perímetro de um círculo.
*/

type Circulo struct{
	raio float64
}

func (c Circulo) calculaArea() float64{
	area := math.Pi * math.Pow(c.raio, 2)
	return area
}

func (c Circulo) calculaPerimetro() float64{
	perimetro := 2 * math.Pi * c.raio
	return perimetro
}

func main(){
	bolinha := Circulo{
		raio: 5,
	}
	fmt.Printf("A area é %.2f e o perímetro é %.2f.\n", bolinha.calculaArea(), bolinha.calculaPerimetro())
}
