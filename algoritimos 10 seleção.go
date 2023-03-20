package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Idade:")
	fmt.Scanln(&idade)
	switch {
	case idade <= 9 && idade >= 0:
		fmt.Println("classificação: Mirim")
	case idade >= 10 && idade < 14:
		fmt.Println("classificação: Infantil")
	case idade >= 14 && idade < 18:
		fmt.Println("classificação: Juvenil")
	case idade >= 18:
		fmt.Println("classificação: Adulto")
	default:
		fmt.Println("Nenhuma das classificações")

	}
}
