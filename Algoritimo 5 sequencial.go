package main

import "fmt"

func main() {
	var idade int
	fmt.Println("idade em anos:")
	fmt.Scanln(&idade)
	var Idade_em_dias int
	Idade_em_dias = idade * 365
	fmt.Println("A Idade em dias é", Idade_em_dias)

}
