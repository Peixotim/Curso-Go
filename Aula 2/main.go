package main

import "fmt"

// Variaveis mudaveis

// Pode ser definido passando o tipo do numero
var num1 int = 4
var num2 int = 4

// Pode ser definido nao passado o tipo do numero
var num3 = 4

//Variaveis constantes / Imutaveis

const name string = "Pedro"

func main() {

	// Pode ser definido com := , mas apenas dentro do escopo da funcao
	num4 := 4

	var sum int = num1 + num2
	var sumTree int = num1 + num2 + num3
	var sumFour int = num1 + num2 + num3 + num4

	const goodMorning string = "Good Morning " + name
	fmt.Println(sum)
	fmt.Println(sumTree)
	fmt.Println(sumFour)
	fmt.Println(goodMorning)
}
