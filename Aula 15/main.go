package main

import (
	"fmt"
	"strings"
)

func main() {

	multiplica := func(x int) int {
		return x * x
	}
	fmt.Println(multiplica(4))
	fmt.Println(soma(20, 02))
	fmt.Println(Uper("olá"))
}

func soma(a, b int) int {
	return a + b
}

func Uper(text string) string {
	return strings.ToUpper(text)
}

func Lower(text string) string {
	return strings.ToLower(text)
}

func Compare(a, b string) bool {
	return strings.EqualFold(a, b)
}
