package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println("Número original:", numero)//20
	fmt.Println("Número invertido:", numeroInvertido)//-20
	//o numero original nao foi invertido, pois o go passa os valores por copia, ou seja, ele cria uma copia do valor na memoria e passa essa copia para a funcao, assim o valor original nao é alterado. Para alterar o valor original devemos passar o endereço de memoria do valor original, ou seja, passar o ponteiro do valor original.

	numero2 := 30
	inverteSinalPonteiro(&numero2)//passando o endereço de memoria do numero2, ou seja, o ponteiro do numero2
	fmt.Println("Número original:", numero2)//-30
}