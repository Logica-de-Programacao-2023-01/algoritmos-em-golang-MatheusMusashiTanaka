package main

import "fmt"

func main() {
	soma := 0.0
	var lista = []float64{1, 2, 3, 4, 5, 6}
	for _, x := range lista {
		soma += x
	}
	media := 0.0
	media = soma / float64(len(lista))
	fmt.Println(media)
}







package main

import "fmt"

func main() {
	var matriz [3][4]int
	for linha := range matriz {
		for coluna := range matriz[linha] {
			matriz[linha][coluna] = linha + coluna
		}
	}
	fmt.Println(matriz)
}
