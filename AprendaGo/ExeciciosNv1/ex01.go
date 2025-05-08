package main

import "fmt"

func main() {
	/*
			Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
		    	1. 42
		    	2. "James Bond"
		    	3. true
			Agora demonstre os valores nestas variáveis, com:
		    	1. Uma única declaração print
		    	2. Múltiplas declarações print
			Solução: https://play.golang.org/p/yYXnWXIQNa
	*/

	x, y, z := 42, "James Bond", true

	fmt.Print(x, y, z)
	fmt.Print(x)
	fmt.Print(y)
	fmt.Print(z)
}
