package main

import "fmt"

const (
	_ = iota + 2025
	year1 = iota + 2025
	year2 = iota + 2025
	year3 = iota + 2025
	year4 = iota + 2025
)

func main() {
	/* 
	Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
	Demonstre estes valores.
	Solução: https://play.golang.org/p/zRBEooRvo4 
	
	*/

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}