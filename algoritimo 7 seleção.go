package main

import "fmt"

func main() {
	var salario float64
	fmt.Println("salario do funcionario em reais:")
	fmt.Scanln(&salario)
	if salario < 1000 {
		salario_com_aumento := salario * 1.1
		fmt.Println("o salario com aumento sera:", salario_com_aumento)
	} else if salario >= 1000 {
		salario_com_aumento_2 := salario * 1.05
		fmt.Println("o salario com aumento sera", salario_com_aumento_2)
	}

}
