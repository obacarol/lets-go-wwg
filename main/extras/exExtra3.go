package main

import (
	"fmt"
)

func main() {
	oceania := make(map[string]string)
	frequencia := make(map[string]int64)

	oceania["Sydney"] = "Australia"
	oceania["Melbourne"] = "Australia"
	oceania["Brisbane"] = "Australia"
	oceania["Camberra"] = "Australia"
	oceania["Adelaide"] = "Australia"
	oceania["Auckland"] = "Nova Zelandia"
	oceania["Wellington"] = "Nova Zelandia"
	oceania["Queenstown"] = "Nova Zelandia"
	oceania["Picton"] = "Nova Zelandia"
	oceania["Suva"] = "Fiji"
	oceania["Nadi"] = "Fiji"
	oceania["Ba"] = "Fiji"
	oceania["Nacualofa"] = "Tonga"
	oceania["Neiafu"] = "Tonga"

	frequencia["Australia"] = 0
	frequencia["Nova Zelandia"] = 0
	frequencia["Fiji"] = 0
	frequencia["Tonga"] = 0

	for k, _ := range frequencia {
		for _, v := range oceania {
			if k == v {
				frequencia[k]++
			}
		}
	}
	fmt.Println(frequencia)
}