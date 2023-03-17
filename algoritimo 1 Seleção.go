package main

import "fmt"

func main() {
	var numero_1 float64
	fmt.Println("primeiro numero:")
	fmt.Scanln(&numero_1)
	var numero_2 float64
	fmt.Println("segundo numero:")
	fmt.Scanln(&numero_2)

	if numero_1 > numero_2 {
		fmt.Println("O primeiro numero é maior que o segundo")
	} else if numero_1 == numero_2 {
		fmt.Println("Os numeros sao iguais")
	} else if numero_1 < numero_2 {
		fmt.Println("O segundo numero é maior que o primeiro")
	}
}

