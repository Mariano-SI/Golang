package main

import "fmt"

func main() {
	//Print, Println, Printf = Usados para exibir diretamente na tela:
	fmt.Print("Mariano")              // Imprime sem quebra de linha
	fmt.Println("Mariano")            // Imprime com quebra de linha ao final
	fmt.Printf("Olá, %s!", "Mariano") // Imprime com formatação (substitui %s por "Mariano")

	/* 
	Principais verbos de formatação:
	%s = string

	%d = inteiro decimal

	%f = número decimal (float)

	%v = valor "padrão" (serve para qualquer tipo)

	%+v = imprime struct com nomes dos campos

	%#v = representa a sintaxe do Go
	*/

	//Sprint, Sprintln, Sprintf = Usados para formatar valores como string (não imprimem na tela, apenas retornam a string):
	y := fmt.Sprint("Teste", " ", "Mariano")
	fmt.Println(y) // Teste Mariano

	z := fmt.Sprintln("Teste", " ", "Mariano")
	fmt.Println(z) // Teste  Mariano\n (com quebra de linha ao final)

	x := fmt.Sprintf("Olá, %s! Você tem %d anos.", "Mariano", 25)
	fmt.Println(x) // Olá, Mariano! Você tem 25 anos.

}
