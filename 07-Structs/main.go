package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco //struct dentro de outra struct	
}

type endereco struct {
	rua    string	
	numero int
}

func main() {
	//struct é o que mais se aproxima de uma classe em go, mas não é uma classe. É um tipo de dado que pode conter outros tipos de dados. É um tipo de dado composto.

	//criando um usuario com var
	var usuario1 usuario
	usuario1.nome = "Mariano"
	usuario1.idade = 20
	fmt.Println(usuario1)
	println(usuario1.idade)

	var endereco1 = endereco{"Rua 1", 123}

	//criando um usuario com :=
	usuario2 := usuario{"Paola", 20, endereco1}
	println(usuario2.nome)

	//criando usuario com var e atribuindo valores
	usuario3 := usuario{"Lucas", 20, endereco1}
	println(usuario3.nome)
	println(usuario3.idade)

	//criadno usuario com new
	usuario4 := new(usuario) //cria um ponteiro para o usuario
	usuario4.nome = "Pedro"
	usuario4.idade = 20
	println(usuario4.nome)
	println(usuario4.idade)

	//criando usuario com const
	//const usuario4 usuario = usuario{"Lucas", 20} //não é possivel criar uma struct com const, pois const é para tipos primitivos e não compostos.
}