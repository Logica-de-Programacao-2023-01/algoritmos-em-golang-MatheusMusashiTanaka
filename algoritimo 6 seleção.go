package main

import "fmt"

func main() {
	var numero_1 int
	fmt.Println("numero 1:")
	fmt.Scanln(&numero_1)
	var numero_2 int
	fmt.Println("Numero 2:")
	fmt.Scanln(&numero_2)
	if numero_1 > 0 && numero_2 > 0 {
		multiplicados := numero_1 * numero_2
		fmt.Println("Os dois numeros multiplicados é:", multiplicados)
	} else {
		somados := numero_1 + numero_2
		fmt.Println("Os dois numeros somados são:", somados)
	}
}
