package main

import "fmt"

func main()  {
	cores := map[string]string{
		"preto" : "#000000",
		"azul" : "#0000FF",
		"cyan" : "#00FFFF",
		"verde" : "#008000",
		"roxo" : "#A020F0",
	}
	fmt.Println(cores)

	delete(cores, "preto")
	fmt.Println(cores)
}
