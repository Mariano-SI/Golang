package main

import "fmt"

func main() {
	x := 10

	if x > 100 {
		fmt.Println("É menor que 100")
	}

	// if init

	if y := 1; y == 1{
		fmt.Println("y é igual a 1")
	}

	fmt.Println(y) //erro pois ele só é acessivel no scopo do init e depois "morre"
}