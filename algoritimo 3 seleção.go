1º modo -> Apenas numeros inteiros
--------------------------------------------------------------------------------------
package main

import "fmt"

func main() {
	var numero_1 int
	fmt.Println("numero:")
	fmt.Scanln(&numero_1)
	if (numero_1 % 2) == 0 {
		fmt.Println("o numero é par")
	} else {
		fmt.Println("o numero é impar")
	}
}
______________________________________________________________________________________
2º modo -> numeros decimais

package main

import "fmt"

func main() {
	var numero float64
	fmt.Println("numero:")
	fmt.Scanln(&numero)
	if (numero*100)%2 == 0 {
		fmt.Println("o numero é par")
	} else {
		fmt.Println("o numero é impar")
	}
}

error mesage : invalid operation: operator % not defined on (numero * 100) (value of type float64)

????????????????????????????
