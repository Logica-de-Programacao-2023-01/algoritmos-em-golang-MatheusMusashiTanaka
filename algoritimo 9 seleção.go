package main

import "fmt"

func main() {
	var numero_1 int
	fmt.Println("primeiro numero:")
	fmt.Scanln(&numero_1)
	var numero_2 int
	fmt.Println("segundo numero:")
	fmt.Scanln(&numero_2)
	var numero_3 int
	fmt.Println("terceiro numero")
	fmt.Scanln(&numero_3)
	if numero_1 > numero_2 && numero_1 > numero_3 {
		if numero_2 > numero_3 {
			fmt.Printf("%d,%d,%d", numero_1, numero_2, numero_3)
		} else {
			fmt.Printf("%d,%d,%d", numero_1, numero_3, numero_2)
		}
	} else if numero_2 > numero_1 && numero_2 > numero_3 {
		if numero_1 > numero_3 {
			fmt.Printf("%d,%d,%d", numero_2, numero_1, numero_3)
		} else {
			fmt.Printf("%d,%d,%d", numero_2, numero_3, numero_1)
		}
	} else if numero_3 > numero_1 && numero_3 > numero_2 {
		if numero_2 > numero_1 {
			fmt.Printf("%d,%d,%d", numero_3, numero_2, numero_1)
		} else {
			fmt.Printf("%d,%d,%d", numero_3, numero_1, numero_2)
		}
	} else if numero_1 == numero_2 && numero_2 == numero_1 {
		fmt.Println("todos os numeros sao iguais!")
	}
}
