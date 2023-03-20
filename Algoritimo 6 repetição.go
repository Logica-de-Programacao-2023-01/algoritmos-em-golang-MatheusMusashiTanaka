package main

import "fmt"

func main() {
	var i int
	fmt.Println("numero:")
	fmt.Scanln(&i)
	for multiplicação := 0; multiplicação < 11; multiplicação++ {
		fmt.Println("O numero multiplicado por", multiplicação, "é", multiplicação*i)
	}
}
