package main

import "fmt"

func main() {
	var peso float64
	fmt.Println("Peso em quilos")
	fmt.Scanln(&peso)
	var altura float64
	fmt.Println("Altura em metros:")
	fmt.Scanln(&altura)

	var IMC float64
	IMC = peso / (altura * altura)
	fmt.Println("IMC:", IMC)

}
