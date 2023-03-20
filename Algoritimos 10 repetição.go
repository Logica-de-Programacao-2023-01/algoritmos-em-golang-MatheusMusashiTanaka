package main

import "fmt"

func main() {
	var numero, maior int
	maior = 0
	for {
		fmt.Println("numero:")
		fmt.Scanln(&numero)
		if numero == 0 {
			break
		}
		if numero > maior {
			maior = numero
		}
	}
	fmt.Println("o maior numero é:", maior)
}
