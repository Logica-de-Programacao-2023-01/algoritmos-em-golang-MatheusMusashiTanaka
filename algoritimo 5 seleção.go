package main

import "fmt"

func main() {
	var numero_1 int
	fmt.Println("numero:")
	fmt.Scanln(&numero_1)
	if (numero_1%3) == 0 && (numero_1%5) == 0 {
		fmt.Println("O numero é multiplo de 3 e 5")
	} else {
		fmt.Println("o numero não é multiplo de 3 e 5")
	}
}
