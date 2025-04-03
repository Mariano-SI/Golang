package main

import "fmt"

func main() {
	//estruturas de controle servem para alterar o funcionamento do codigo a partir de consiçoes

	var numero = 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	}else if numero == 15 {
		fmt.Println("Igual a 15")
	}else{
		fmt.Println("Menor que 15")
	}

	//if init

	if outroNumero := numero; outroNumero >= 0 { //ja declara e faz uma condição
		fmt.Println("Numero é maior que zero")
	}else{
		fmt.Println("É menor que 0")
	}

	//fmt.Println(outroNumero) // aqui nao consigo acessar aquela variavel pois seu escopo é apenas no if

}