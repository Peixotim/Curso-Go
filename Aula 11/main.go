package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Quando é quarta ?")
	today := time.Now().Weekday()

	switch time.Monday {
	case today + 0:
		fmt.Println("Hoje é Segunda")
	case today + 1:
		fmt.Println("Hoje é Terca Feira")
	case today + 2:
		fmt.Println("Hoje é Quarta Feira !")
	case today + 3:
		fmt.Println("Hoje é Quinta Feira !")
	case today + 4:
		fmt.Println("Hoje é Sexta Feira !")
	case today + 5:
		fmt.Println("Hoje é Sabado !")
	case today + 6:
		fmt.Println("Hoje é Domingo !")
	}
}
