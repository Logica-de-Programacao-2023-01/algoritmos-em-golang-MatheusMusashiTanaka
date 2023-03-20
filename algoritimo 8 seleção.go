package main

import "fmt"

func main() {
	var numero_do_dia int
	fmt.Println("o dia (em numero) é :")
	fmt.Scanln(&numero_do_dia)
	switch numero_do_dia {
	case 1:
		fmt.Println("Domingo!!")
	case 2:
		fmt.Println("Segunda !!")
	case 3:
		fmt.Println("Terça!!")
	case 4:
		fmt.Println("Quarta!!")
	case 5:
		fmt.Println("Quinta!!")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("sabado")
	default:
		fmt.Println("nao a esse numero")
	}
}
