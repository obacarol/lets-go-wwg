package main

import "fmt"

/*
Realize a implementação de uma pilha e seus dois principais métodos, push() e pop().
*/

type Pilha []string

func main() {
	var pedidos Pilha

	pedidos.Push("Pizza 4 queijos")
	pedidos.Push("Pizza Portuguesa")
	pedidos.Push("Pizza de Calabresa")

	fmt.Println("A pilha de pedidos é: ", pedidos)

	for len(pedidos) > 0 {
		x, y := pedidos.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}

//isEmpty - verifica se a pilha está vazia
func (p *Pilha) IsEmpty() bool{
	return len(*p) == 0
}

//Push - insere uma string nova no final da pilha
func (p *Pilha) Push(pedido string) {
	*p = append(*p, pedido)
}

//Pop - remove e retorna a string do topo da pilha
func (p *Pilha) Pop() (string, bool) {
	if p.IsEmpty() {
		return "", false
	} else {
		i := len(*p) - 1
		elemento := (*p)[i]
		*p = (*p)[:i]
		return elemento, true
	}
}