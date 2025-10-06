package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello string = "Olá Mundo !"

	var question string = "Como vai?"

	var meet string = hello + question
	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet))
	fmt.Println(strings.ToLower(meet))
	fmt.Println(strings.Count(meet, "o"))        // Conta quantas strings do mesmo tipo ou igual tem
	fmt.Println(strings.Contains(meet, "mundo")) //Verifica se possui a palavra
}
