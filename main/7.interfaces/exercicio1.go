package main

import (
	"fmt"
	"math"
)

/*
Crie tipos para representar quadrados e círculos.
Crie uma interface que descreve o comportamento de calcular a área de
	uma forma geométrica com a seguinte assinatura: calculeArea() float64.
Implemente esse comportamento para os dois tipos criados.
Depois, crie uma função que tem como parâmetro a interface que você criou e
	que imprime o relatório do cálculo da área da forma geométrica.
Demonstre que seus tipos implementam a interface que você criou passando
	valores desses tipos como argumentos na chamada dessa função.
*/

type Quadrado struct{
	lado float64
}

type Circulo struct{
	raio float64
}

type formaGeometrica interface{
	calculeArea() float64
}

func (q Quadrado) calculeArea() float64{
	return math.Pow(q.lado, 2)
}

func (c Circulo) calculeArea() float64{
	area := math.Pi * math.Pow(c.raio, 2)
	return area
}

func imprimirArea(f formaGeometrica) {
	fmt.Printf("A área é %.2f m².\n", f.calculeArea())
}

func main(){
	quadrado := Quadrado{
		lado: 4,
	}

	circulo := Circulo{
		raio: 2.5,
	}

	imprimirArea(quadrado)
	imprimirArea(circulo)
}