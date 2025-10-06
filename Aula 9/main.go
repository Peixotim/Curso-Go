package main

import "fmt"

// Obervacao map é uma estrutura chave valor
// E sua chave deve ser unica
func main() {
	pessoas := map[string]int64{}
	pessoas["Pedro"] = 18
	pessoas["Maria"] = 19
	pessoas["Pokemon"] = 2
	fmt.Println(pessoas)
	fmt.Println("A idade de Pedro é ", pessoas["Pedro"])

	//Verificando se há no map a pessoa pesquisada
	if idade, ok := pessoas["Pokemon"]; ok {
		fmt.Println("Pessoa existe no map", idade, ok)
	} else {
		fmt.Println("Pessoa nao existe no map")
	}

	//Funcao delete ou seja tirar uma chave do nosso map
	//Passa o map e depois o nome da chave para ser retirada
	delete(pessoas, "Pokemon")

	fmt.Println(pessoas)
}
