package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	Pedro := Pessoa{
		Nome:  "Pedro",
		Idade: 18,
	}

	user := &Pedro
	fmt.Println(Pedro)
	fmt.Println(user)
}
