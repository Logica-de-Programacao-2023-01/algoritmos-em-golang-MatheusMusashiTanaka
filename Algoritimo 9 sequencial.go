package main

import "fmt"

func main() {
	var preço float64
	fmt.Println("preço do produto em reais:")
	fmt.Scanln(&preço)
	var desconto float64
	desconto = preço * 0.10
	var produto_com_desconto float64
	produto_com_desconto = preço - desconto
	fmt.Println("O novo preço é", produto_com_desconto, "reais")
}
