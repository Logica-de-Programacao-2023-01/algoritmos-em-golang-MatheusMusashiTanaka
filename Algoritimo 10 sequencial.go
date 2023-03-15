package main

import "fmt"

func main() {
	var peso float64
	fmt.Println("peso em quilos:")
	fmt.Scanln(&peso)
	var convertor float64
	convertor = peso * 2.205
	fmt.Println("O peso em libras é", convertor)
}

