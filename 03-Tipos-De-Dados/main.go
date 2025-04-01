package main

import (
	"errors"
	"fmt"
)

func main() {
	//tipos numericos
	//inteiros: int, int8, int16, int32, int64 = cada numero representa a quantidade de bits. O int sem especificação é o mesmo que o int32 ou int64 dependendo do sistema operacional (32 ou 64 bits)

	var numero int64 = 100000000000
	fmt.Println(numero)
	//inteiros sem sinal: uint, uint8, uint16, uint32, uint64 = cada numero representa a quantidade de bits
	// numeros reais: float32, float64 = cada numero representa a quantidade de bits. O float sem especificação é o mesmo que o float32 ou float64 dependendo do sistema operacional (32 ou 64 bits)
	var numeroReal float64 = 10.5
	fmt.Println(numeroReal)

	//strings: string = conjunto de caracteres.
	var nome string = "Lucas"
	nome2 := "Mariano"
	fmt.Println(nome2)
	fmt.Println(nome)
	
	//booleanos: bool = verdadeiro ou falso.
	var verdadeiro bool = true
	fmt.Println(verdadeiro)
	falso := false
	fmt.Println(falso)

	//tipo errro: error = representa um erro. O tipo error é uma interface que possui um método Error() que retorna uma string com a descrição do erro.
	var erro error = errors.New("Erro de teste")	
	fmt.Println(erro)

}