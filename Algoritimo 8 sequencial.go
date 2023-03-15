package main

import "fmt"

func main() {
	var dias int
	fmt.Println("dias trabalhados:")
	fmt.Scanln(&dias)
	var diaria int
	fmt.Println("Valor diaria")
	fmt.Scanln(&diaria)
	var salario int
	salario = (dias * diaria)
	fmt.Println("O salario é", salario, "Reais")

}

