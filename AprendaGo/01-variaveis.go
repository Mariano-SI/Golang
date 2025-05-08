package main

import "fmt"

func main() {

	// := infere o tipo, se comporta como var, permitindo reatribuição e deve ser inicializado, so funciona entre chaves, nao funciona no escopo de pacote
	nome := "Mariano"
	nome = "outro"

	fmt.Println(nome) //outro


	x := 10
	y := "olá bom dia"

	fmt.Printf("x: %v, %T \n", x,x)
	fmt.Printf("Y: %v, %T", y,y)

	// var pode ou nao ser inicializada de cara, se for incializada ela consegue inferir o tipo, se nao, devemos lhe explicitar. Var aceita reatribuição

	var teste string //nao inicializei = tive que explicitar o tipo

	var teste2 = "mariano" //inicializei = nao preciso (mas posso) explicitar o tipo

	fmt.Println(teste, teste2)



}
