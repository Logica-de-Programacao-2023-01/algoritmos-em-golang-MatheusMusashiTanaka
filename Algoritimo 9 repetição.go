package main

import "fmt"

func main() {
	var numero int
	lista := []int{}
	soma := 0
	for {
		fmt.Println("numero:")
		fmt.Scanln(&numero)
		if numero == 0 {
			break
		}
		lista = append(lista, numero)
		soma += numero
	}
	media := soma / len(lista)
	fmt.Println("a media aritimetica é", media)
}
