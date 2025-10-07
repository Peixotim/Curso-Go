package main

import (
	"fmt"
	"time"
)

func exibirMensagem() {
	fmt.Println("Olá isto é uma goroutine !")
}

func main() {
	go exibirMensagem()
	time.Sleep(1 * time.Second)
	fmt.Println("Olá main function!")
}
