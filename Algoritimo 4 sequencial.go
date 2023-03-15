package main

import "fmt"

func main() {
	var numero_1 float64
	fmt.Println("Primeiro numero:")
	fmt.Scanln(&numero_1)
	var numero_2 float64
	fmt.Println("Segundo numero:")
	fmt.Scanln(&numero_2)
	var numero_3 float64
	fmt.Println("Terceiro numero:")
	fmt.Scanln(&numero_3)
	var media float64
	media = ((numero_1 * 2) + (numero_2 * 3) + (numero_3 * 5)) / 10
	fmt.Println("A media ponderada (peso 2, 3 e 5 respectivamente) é", media)

}
