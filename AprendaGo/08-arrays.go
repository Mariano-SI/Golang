package main

import "fmt"

func main() {
	//array

	var x [5]int
	x[0] = 1
	x[1] = 10
	x[2] = 2
	x[4] = 6

	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(len(x))

	//slices

	var meuSlice = []string{"Mariano", "teste"}

	meuSlice = append(meuSlice, "nova string")

	fmt.Println(meuSlice)
	fmt.Println(len(meuSlice))

	//funcao range

	frutas := []string{"banana", "maca", "laranja"}

	for index, value := range frutas {
		fmt.Println(index, value)
	}

	//fatiar um slice

	sabores := []string{"pepperoni", "mussarela", "abacaxi", "quatro queijos", "marguerita"}
	//quero pegar os 2 primeiros elementos 
	fatia := sabores[0:2] //a partir do indice zero ate o indice 2 onde o indice 2 nao esta incluso

	//quero pegar "mussarela", "abacaxi", "quatro queijos",

	fatia2 := sabores[1:4]

	//outros exemplos da sintaxe [:]
	// [x:len()] do indice x ate o final 
	// [:y] do indice y ate o inicio

	fmt.Println("Fatia", fatia)
	fmt.Println("Fatia2", fatia2)

	//deletar

	//deletar abacaxi

	sabores = append(sabores[:2], sabores[3:]...) //seleciona tudo que tem antes e tudo que tem depois do item a ser excluido

	//funcao make: serve para deixar o uso de slices mais performatico colocando uma capacidade inciial pra evitar a recriação do slice sempre que a capacidade é estourada. é util quando sabemos o numero de elementos ou apenas sabemos que o nuero é alto e queremos evitar a recriação

	sliceComMake := make([]int, 5, 10)

	sliceComMake[0], sliceComMake[1], sliceComMake[2], sliceComMake[3], sliceComMake[4] = 1,2,3,4, 5

	sliceComMake = append(sliceComMake, 6)
	sliceComMake = append(sliceComMake, 7)
	sliceComMake = append(sliceComMake, 8)
	sliceComMake = append(sliceComMake, 9)
	sliceComMake = append(sliceComMake, 10)

	fmt.Println("sliceComMake", sliceComMake, len(sliceComMake), cap(sliceComMake))

	sliceComMake = append(sliceComMake, 10) //passar a capacidade

	fmt.Println("sliceComMake", sliceComMake, len(sliceComMake), cap(sliceComMake)) //cap = 20

	//slices multi dimencionais = slice de slices

	sliceDeSlice := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}

	fmt.Println(sliceDeSlice)

	//map: sao excelentes em operacoes de busca

	amigos := map[string]int{
		"Alfredo":33333,
		"Joaozinho": 123414,
	}

	amigos["Teste"] = 124214


	fmt.Println(amigos)
	fmt.Println(amigos["Joaozinho"])

	for key, value := range amigos {
		fmt.Println(key,value)
	}

	//como deletar um elemento do map
	delete(amigos, "Joaozinho")

	fmt.Println(amigos)
}