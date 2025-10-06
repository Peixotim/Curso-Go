package main

import "fmt"

//Tipos de inteiro

// - **int 8 ** Inteiro de 8 bits (valores de -128 a 127)
// - **int 16 ** Inteiro de 16 bits (valores de -32768 a 32767)
// - **int 32 ** Inteiro de 32 bits (valores de -2137383648 a 2147483647)
// - **int 64 ** Inteiro de 64 bits (valores de -9223372036854775808 a 9223372036854775807)
func main() {

	var numberOne int8 = 8
	var numberTwo int16 = 16
	var numberTree int32 = 32
	var numberFour int64 = 64

	fmt.Println(numberOne)
	fmt.Println(numberTwo)
	fmt.Println(numberTree)
	fmt.Println(numberFour)

}
