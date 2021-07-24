package main

import "fmt"

/*
Considerando os times do item anterior, crie uma slice para representar cada um.
Adicione o jogador Luis Inácio no time vermelho usando a função append()
Printe na tela os nomes dos jogadores do time vermelho.
*/

func main()  {
	timeAmarelo := []string{"Fernando", "João", "Lucia", "Mariana", "Ana"}
	timeVermelho := []string{"Helena", "Jonas", "José", "Juliana"}

	timeVermelho = append(timeVermelho, "Luis Inacio")

	fmt.Println("Os jogadores do time amarelo são", timeAmarelo)
	fmt.Println("Os jogadores do time vermelho são", timeVermelho)
}
