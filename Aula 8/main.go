package main

import "fmt"

// Funcao Slices
func main() {
	var pokedex []string

	pokedex = append(pokedex, "Pikachu", "Charizard", "RayQuaza", "Charmander", "Lucario", "Slot Vazio")
	//A funcao len é utilizada para verificar o tamanho de algumas estruturas em go
	//Para imprimir so um certo indice é feito assim slice[x:x-1]
	//No caso dentro do colchetes é o seguinte [indiceDesejado : ateOIndiceDesejado - 1]
	fmt.Println(len(pokedex))
	//Ele pega o primeiro indice e o terceiro indice menos 1 ou seja pega o segundo indice
	fmt.Println(pokedex[0:2])
	//Podemos definir tambem um inicio e ele rodar todos
	fmt.Println(pokedex[0:])
	//Ou se nao definirmos apenas o final
	fmt.Println(pokedex[:4])

	//Agora para removermos o indice e dar um novo indice final
	pokedex = pokedex[:2]
	pokedex = append(pokedex, "Ratata")

	fmt.Println(pokedex)
}
