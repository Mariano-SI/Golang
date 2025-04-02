package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa   //composição, não herança
	matricula string
	curso     string
}

func main() {
	//não existe herança em go

	pessoa1 := pessoa{"Mariano", "silva", 23, 173}

	estudante1 := estudante{pessoa1, "123456", "Sistemas"}
	fmt.Println(estudante1) //acessando o nome da pessoa dentro do estudante

	var estudante2 estudante
	estudante2.nome = "Paola"
	estudante2.sobrenome = "Kunze"
	estudante2.idade = 27
	estudante2.altura = 170
	estudante2.matricula = "123456"
	fmt.Println(estudante2) //acessando o nome da pessoa dentro do estudante


	estudante3 := estudante{pessoa{"Mariano", "silva", 23, 173}, "123456", "Sistemas"}
	
	fmt.Println(estudante3) //acessando o nome da pessoa dentro do estudante

	var estudante4 estudante = estudante{
		pessoa: pessoa{"Lucas", "silva", 23, 173},
		matricula: "123456",
		curso: "Sistemas",
	}

	fmt.Println(estudante4) //acessando o nome da pessoa dentro do estudante
}