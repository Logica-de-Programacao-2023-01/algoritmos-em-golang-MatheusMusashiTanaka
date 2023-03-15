package main

import "fmt"

func main() {
	var numero int
	fmt.Println("numero:")
	fmt.Scanln(&numero)
	var Aumento int
	Aumento = (numero + 1)
	var diminuir int
	diminuir = (numero - 1)
	fmt.Println("O sucessor sera", Aumento)
	fmt.Println("O antecessor", diminuir)

}
