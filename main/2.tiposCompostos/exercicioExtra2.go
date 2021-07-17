package main

import "fmt"

func main()  {
	timeAmarelo := []string{"Fernando", "João", "Lucia", "Mariana", "Ana"}
	timeVermelho := []string{"Helena", "Jonas", "José", "Juliana"}

	timeVermelho = append(timeVermelho, "Luis Inacio")

	fmt.Println("Os jogadores do time amarelo são", timeAmarelo)
	fmt.Println("Os jogadores do time vermelho são", timeVermelho)
}
