package main

import "fmt"

func main() {
	var numero_1 int
	fmt.Println("numero:")
	fmt.Scanln(&numero_1)
	if (numero_1 % 2) == 0 {
		fmt.Println("o numero é par")
	} else {
		fmt.Println("o numero é impar")
	}
}
