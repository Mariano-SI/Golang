package main

import "fmt"

func main() {
	//arrays é uma lista de valores de um mesmo tipo. Ele tem tamanho fixo, pode ser incializado já com valores ou vazio. sempre dwevemos especificar o tamanho do array em sua inciialização, caso constrario ele se tornara um slice

	var array1 [5]int //array de 5 inteiros
	array1[0] = 1
	array1[4] = 2

	fmt.Println(array1, len(array1)) //imprime o array inteiro

	var array2 [5]int = [5]int{1, 2, 3, 4, 5} //array de 5 inteiros
	fmt.Println(array2) //imprime o array inteiro

	array3 := [5]int{1, 2,} //array 
	fmt.Println(array3) //imprime o array inteiro


	//slice é uma estrutura baseada no array porem com flexibilidades a mais
	//é mais comum usar slice do que array, pois o slice é mais flexivel e mais facil de usar
	//ele não tem tamanho fixo, e pode ser redimensionado, ou seja, podemos adicionar e remover elementos dele

	var slice1 []int //slice de inteiros, não tem tamanho fixo
	slice1 = append(slice1, 18) // adiciona ao slice e retorna um novo
	fmt.Println(slice1)


	slice2 := []int{1, 3, 5}
	fmt.Println(slice2)
}