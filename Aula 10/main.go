package main

import (
	"errors"
	"fmt"
)

func main() {
	nota := 10
	nota2 := 7
	nota3 := 6
	notaTotal := (nota + nota2 + nota3) / 3

	if notaTotal >= 7 {
		fmt.Println("Aprovado !")
	} else {
		fmt.Println("Reprovado !")
	}
	fmt.Println(notaTotal)

	if err := thisIsAnError(); err != nil {
		fmt.Println(err.Error())
	}

	players := map[string]int64{
		"Pedro": 20,
	}

	fmt.Println(players)
}

func thisIsAnError() error {
	return errors.New("Isto é um erro")
}
