package main

import "fmt"

type user struct {
	name string
	password string
}

func main() {
	//ponteiro em programação é uma variavel que guarda um endereço de memoria de alguma coisa

	var variavel1 int = 10
	var variavel2 int = variavel1 //atribuições em go por padrao sao copias, indenpendente do tipo da variavel. Diferente de javascript e outras linguagens onde apenas tipo primitivos seriam copia e os tipos tipo referencia sao passados por referencia

	fmt.Println(variavel1, variavel2) //10 10

	variavel1++

	fmt.Println(variavel1, variavel2)// 11 10


	var user1 user = user{"Mariano", "123"}
	var user2 = user1

	fmt.Println(user1, user2)

	user1.name = "Joãozinho"

	fmt.Println(user1, user2) // a mudança em user1 nao reflete em user2 pois sao copias nao referencias

	//PONTEIRO = REFRENCIA DE MEMORIA

	var variavel3 int = 100
	var ponteiro *int = &variavel3 //a variavel 3 é um ponteiro de int (*int) e recebeu o enredeço de variavel 3 (&variavel3)

	fmt.Println(variavel3, ponteiro)// 100 0xc00000a130

	//como vimos ponteiro fguarda um endereço e para vermos esse endereço usamos * antes da variavel

	fmt.Println(*ponteiro) //100 desreferenciação

	*ponteiro++

	fmt.Println(variavel3)//101

}