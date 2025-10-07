package main

import (
	"fmt"
)

type Pessoa struct {
	Nome   string
	Idade  int
	Genero string
}

type ContaBancaria struct {
	Pessoa    Pessoa
	Saldo     float64
	TipoConta string
}

func (Conta ContaBancaria) Depositar(Adicionar float64) {
	Conta.Saldo += float64(Adicionar)
}

func (Conta ContaBancaria) Sacar(Sacar float64) {
	Conta.Saldo -= float64(Sacar)
}

func (p Pessoa) Apresentar() {
	fmt.Println("Olá meu nome é ", p.Nome, " e eu tenho ", p.Idade, " anos")
}

func (p Pessoa) CriarConta(TipoConta string) {
	ContaBancaria := ContaBancaria{
		Pessoa:    p,
		Saldo:     0,
		TipoConta: TipoConta,
	}

	fmt.Println("Conta Criada", ContaBancaria)
}

func main() {
	Pedro := Pessoa{
		Nome:   "Pedro",
		Idade:  18,
		Genero: "Masculino",
	}

	Pedro.CriarConta("Corrente")
}
