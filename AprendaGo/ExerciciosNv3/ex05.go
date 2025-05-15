package main

import "fmt"

func main() {
	//Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

	for i := 10; i <= 100; i++ {
		restoDaDivisaoPor4 := i % 4
		fmt.Println(restoDaDivisaoPor4)

		if(restoDaDivisaoPor4 == 0){
			fmt.Printf("O numero %v é divisível por 4", i)
		}
	}
}