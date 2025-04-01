package main

import "fmt"

func main() {
	var nome = "Mariano1"
	var nome2 string = "Mariano2"
	var nome3, sobrenome = "Mariano3", "Silva"
	nome4 := "Mariano4" // Declarando e inicializando uma variável com o operador de atribuição curta (:=)

	fmt.Println(nome, nome2, nome3, nome4, sobrenome)

	var num1 = 10
	var num2 = 20

	num1, num2 = num2, num1 // Troca os valores de num1 e num2
	fmt.Println(num1, num2)
}