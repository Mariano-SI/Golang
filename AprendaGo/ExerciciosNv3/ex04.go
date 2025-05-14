package main

import "fmt"

func main() {
	//Crie um loop utilizando a sintaxe: for {} Utilize-o para demonstrar os anos desde que você nasceu.

	anoQueNasci := 2001
	anoAtual := 2025

	for {
		if anoQueNasci == anoAtual {
			break
		}
		fmt.Println(anoQueNasci)
		anoQueNasci++
	}
}