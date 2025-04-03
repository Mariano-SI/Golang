package main

import "fmt"

func main() {
	//go so possui a estrutura de repetição for e podemos lhe usar de diversas maneiras

	//for como while

	i := 0

	for i < 10 {
		fmt.Println("i", i)
		i++
	}
	//for tradicional
	for j := 0; j <10; j++ {
		fmt.Println("j ", j)
	}

	//for como while true

	// k := 0
	// for{
	// 	fmt.Println("k ", k)
	// 	k++
	// }

	//for com range usamos quando vamos iterar algo como array, slice, string, map
	nomes := [3]string{"Mariano", "Paola", "Teste"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}
	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))// se nao converter recebemos o valor ascii de cada letra
	}

	usuario := map[string]string{
		"Nome": "Mariano",
		"Sobrenome": "Silva",
	}

	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}

	//nao da pra fazer range em struct... para isso...

}