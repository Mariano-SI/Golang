package main

import "fmt"

func main() {
	//Crie um programa que demonstre o funcionamento da declaração if.

	idade := 24
	if idade < 30 {
		fmt.Println("Ainda nao chegou aos 30")
	}

	toCansado := false
	if toCansado {
		fmt.Println("Vai descansar")
	}else{
		fmt.Println("Continue estudando")
	}
}