package main

import "fmt"

func main() {
	var numero float64
	fmt.Println("numero:")
	fmt.Scanln(&numero)
	if float64(numero*100)%2 == 0 {
		fmt.Println("o numero é par")
	} else {
		fmt.Println("o numero é impar")
	}
}
