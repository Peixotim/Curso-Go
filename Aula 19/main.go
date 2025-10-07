package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
}
