package main

import "fmt"

type myType1 int

var x1 myType1
var y1 int

func main() {
	/*
			Utilizando a solução do exercício anterior:
		    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
		    2. Na função main:
		        1. Isto já deve estar feito:
		            1. Demonstre o valor da variável "x"
		            2. Demonstre o tipo da variável "x"
		            3. Atribua 42 à variável "x" utilizando o operador "="
		            4. Demonstre o valor da variável "x"
		        2. Agora faça tambem:
		            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
		            2. Demonstre o valor de "y"
		            3. Demonstre o tipo de "y"
		Solução: https://play.golang.org/p/uq81T_fasP
	*/

	fmt.Printf("x1: valor: %v tipo: %T\n", x1, x1)

	x1 = 42

	fmt.Printf("x1: valor: %v\n", x1)

	y1 = int(x1)

	fmt.Println(y1)
	fmt.Printf("tipo de y1: %T", y1)
}