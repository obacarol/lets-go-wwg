package main

import "fmt"

func main() {
	a := 11<=10
	b := 10==10
	c := 12!=11
	d := 13<13
	e := 12>11 && 11>10

	fmt.Printf("O tipo de a é %T e seu valor é %v.\n", a, a)
	fmt.Printf("O tipo de b é %T e seu valor é %v.\n", b, b)
	fmt.Printf("O tipo de c é %T e seu valor é %v.\n", c, c)
	fmt.Printf("O tipo de d é %T e seu valor é %v.\n", d, d)
	fmt.Printf("O tipo de e é %T e seu valor é %v.\n", e, e)
}