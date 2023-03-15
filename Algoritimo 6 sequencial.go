package main

import "fmt"

func main() {
	var salario float64
	fmt.Println("Salario atual:")
	fmt.Scanln(&salario)
	var Aumento float64
	Aumento = (salario * 0.15)
	var Aumento_salario float64
	Aumento_salario = salario + Aumento
	fmt.Println("O aumento sera de", Aumento, "Reais")
	fmt.Println("O salario após o aumento sera", Aumento_salario, "Reais")

}

