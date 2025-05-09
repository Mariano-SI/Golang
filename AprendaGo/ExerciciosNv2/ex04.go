package main

import "fmt"

func main() {
	/*
		Crie um programa que:
		Atribua um valor int a uma variável
		Demonstre este valor em decimal, binário e hexadecimal
		Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
		Demonstre esta outra variável em decimal, binário e hexadecimal
	*/

	initialNumber := 200
	fmt.Printf("%v\n%b\n%x\n", initialNumber, initialNumber, initialNumber)

	anotherNumber := initialNumber << 1

	fmt.Printf("%v\n%b\n%x\n", anotherNumber, anotherNumber, anotherNumber)
}