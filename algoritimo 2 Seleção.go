package main

import "fmt"

func main() {
	var numero_1 float64
	fmt.Println("primeiro numero:")
	fmt.Scanln(&numero_1)
	var numero_2 float64
	fmt.Println("segundo numero:")
	fmt.Scanln(&numero_2)
	var numero_3 float64
	fmt.Println("terceiro numero:")
	fmt.Scanln(&numero_3)

	if numero_1 < numero_2 {
		if numero_1 < numero_3 {
			fmt.Println("o", numero_1, "é o menor")
		} else {
			fmt.Println("O", numero_3, "é o menor")
		}
	} else if numero_1 == numero_2 && numero_1 == numero_3 {
		if numero_1 == numero_3 {
			fmt.Println("todos os numeros sao iguais")
		}
	} else if numero_2 < numero_3 {
		fmt.Println("O", numero_2, "é o menor")
	} else {
		fmt.Println("o", numero_3, "é o menor")
	}
}
