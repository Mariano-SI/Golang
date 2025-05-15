package main

import "fmt"

func main() {
	//Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

	esporteFavorito := "futsal"

	switch esporteFavorito {
	case "volei":
		fmt.Println("É volei")
	case "futsal":
		fmt.Println("É futsal")
	case "futebol":
		fmt.Println("É futebol")
	}
}
