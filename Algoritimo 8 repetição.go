package main

import "fmt"

func main() {
	var numero int
	fmt.Println("digite o numero:")
	fmt.Scanln(&numero)
	for i := numero; i > 0; i-- {
		if numero%i == 0 {
			fmt.Println("o numero é divisivel por ", i)
		}
	}
}
