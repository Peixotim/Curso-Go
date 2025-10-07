package main

import "fmt"

// Assim cria um objeto
type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Rua    string
	Cidade string
	Estado string
}

func main() {

	pedro := Cliente{
		Nome:  "Pedro",
		Idade: 18,
		Endereco: Endereco{
			Rua:    "Avenida JK",
			Cidade: "Timoteo",
			Estado: "Minas Gerais",
		},
		Email: "pedropeixotovz@gmail.com",
	}

	fmt.Println(pedro)
	fmt.Println(pedro.Endereco.Rua)
}
