package main

import (
	"fmt"
)

func main() {

	var numberFloatOne float32 = 2.30
	var numberFloatTwo float64 = 2.40

	//Observacao para fazer a soma de dois numeros o tipo tem que ser o mesmo
	const pi float32 = 3.14
	var raio float32 = 2.5
	var area = pi * raio * raio

	fmt.Println(numberFloatOne)
	fmt.Println(numberFloatTwo)

	fmt.Println("A área do raio é de : ", area)
}
