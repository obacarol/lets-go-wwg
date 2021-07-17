package main

import "fmt"

func main() {
	temPelo := true
	ehFofinha := true
	gostaDeColo := true

	ehAMimi := temPelo && ehFofinha && gostaDeColo

	fmt.Println(ehAMimi)
	fmt.Printf("%T", ehAMimi)
}