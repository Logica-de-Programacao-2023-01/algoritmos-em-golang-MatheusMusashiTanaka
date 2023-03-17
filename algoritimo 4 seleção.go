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
	var sexo string
	fmt.Println("selecione seu sexo:")
	fmt.Println("1. Masculino")
	fmt.Println("2. Feminino")
	fmt.Scanln(&sexo)
	switch sexo {
	case "1":
		if IMC <= 18.5 {
			fmt.Println("Indice de massa corporal abaixo da media")
		} else if IMC > 18.5 && IMC < 24.9 {
			fmt.Println("indice de massa normal")
		} else if IMC > 24.9 {
			fmt.Println("Indice de massa corporal acima da media")
		}
	case "2":
		if IMC <= 18.5 {
			fmt.Println("Indice de massa corporal abaixo da media")
		} else if IMC > 18.5 && IMC < 24.9 {
			fmt.Println("indice de massa normal")
		} else if IMC > 24.9 {
			fmt.Println("Indice de massa corporal acima da media")
		}
	}
}
