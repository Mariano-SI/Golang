package main

import "fmt"

func main() {
	//Crie um programa que utilize a declaração switch, sem switch statement especificado.

	x := 10

	switch {
	case x > 30:
		fmt.Println("maior que trinta")
	case x < 30:
		fmt.Println("menor que trinta")
	default:
		fmt.Println("é trinta")
	}
}