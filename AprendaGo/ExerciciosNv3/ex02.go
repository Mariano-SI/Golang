package main

import "fmt"

func main() {
	/*
			Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
			Por exemplo:
		    65
		        U+0041 'A'
		        U+0041 'A'
		        U+0041 'A'
		    66
		        U+0042 'B'
		        U+0042 'B'
		        U+0042 'B'
		    ...e por aí vai.
	*/

	for i := 65; i <= 90; i++ {
		for j:=0; j<=2; j++{
			fmt.Printf("%#U\n", i)
		}
	}
}