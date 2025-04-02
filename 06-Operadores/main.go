package main

import "fmt"

func main() {
	//artimeticos
	soma := 10 + 5          //soma
	subtracao := 10 - 5     //subtracao
	multiplicacao := 10 * 5 //multiplicacao
	divisao := 10 / 5       //divisao
	modulo := 10 % 5        //modulo

	//Observacao nao é possivel fazer operacoes com numeros de tipos diferentes, por exemplo: int8 e int16, é necessario fazer o cast para o mesmo tipo.
	//var numero int8 = 10	
	//var numero2 int16 = 5
	//resultado := numero + int8(numero2) //cast para o mesmo tipo

	fmt.Println("Soma:", soma)
	fmt.Println("Subtracao:", subtracao)
	fmt.Println("Multiplicacao:", multiplicacao)
	fmt.Println("Divisao:", divisao)
	fmt.Println("Modulo:", modulo)

	//atribuição
	//:= é o operador de atribuição curta, ele cria a variavel e atribui o valor a ela. O tipo da variavel é definido pelo valor atribuido a ela.
	//= é o operador de atribuição, ele atribui o valor a variavel ja existente podendo inferir ou nao

	//comparacao
	maior := 10 > 5         //maior
	maiorOuIgual := 10 >= 5 //maior ou igual
	menor := 10 < 5         //menor
	menorOuIgual := 10 <= 5 //menor ou igual
	igual := 10 == 5        //igual
	naoIgual := 10 != 5     //nao igual

	fmt.Println("Maior:", maior)
	fmt.Println("Maior ou igual:", maiorOuIgual)
	fmt.Println("Menor:", menor)
	fmt.Println("Menor ou igual:", menorOuIgual)
	fmt.Println("Igual:", igual)
	fmt.Println("Nao igual:", naoIgual)

	//logicos
	// and: &&
	// or: ||
	// not: !

	//unarios
	numero := 10
	numero++ //incremento em um
	numero+= 5 //incremento
	numero-- //decremento
	numero-= 5 //decremento
	numero*= 5 //multiplicacao
	numero/= 5 //divisao
	numero%= 5 //modulo

	//ternario
	// nao existe em go, mas pode ser simulado com if else

}