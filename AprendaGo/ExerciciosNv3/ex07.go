package main

import "fmt"

func main() {
	idade := 30
	if idade > 30 {
		fmt.Println("Maior que trinta")
	} else if idade < 30 {
		fmt.Println("Menor que trinta")
	}else{
		fmt.Println("É trinta")
	}
}