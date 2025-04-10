package main

import "fmt"

func clojure() func() {
	texto := "Dentro da funcao clojure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao 
}

func main() {
	

	texto := "Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := clojure() 

	funcaoNova()// Dentro da funcao clojure. Go manteve a referencia inicial
}