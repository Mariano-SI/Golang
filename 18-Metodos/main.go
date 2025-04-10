package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func (u usuario) apresentar()  {
	fmt.Println(fmt.Sprintf("Meu nome é %s e tenho %d anos.", u.nome, u.idade))
} //cria metodo associado a struct usuario, o (u usuario) é o receptor do método, ou seja, ele é o objeto que está chamando o método. O nome do receptor pode ser qualquer coisa, mas é comum usar uma letra ou abreviação do tipo da struct.

func (u usuario) maioridade() bool {
	if u.idade >= 18 {
		return true
	}
	return false
}

func(u *usuario) fazerAniversario() { //passamos um ponteiro de usuario pois queremos manipular suas informações e refletir no original
	u.idade++ //nao precisa desreferenciar o ponteiro pois o go faz isso automaticamente
}

func main() {
	//metodos sao bem parecidos com funcoes porem a diferença entre eles é que o método sempre deve estar associado a algo como structs, interfaces, etc.

	usuario1 := usuario{
		nome: "Mariano",
		idade: 17,
	}

	usuario1.apresentar()

	if(usuario1.maioridade()) {
		fmt.Println("O usuário é maior de idade.")
	}else{
		fmt.Println("O usuário é menor de idade.")
	}

	usuario1.fazerAniversario() //fazendo aniversario do usuario1

	if(usuario1.maioridade()) {
		fmt.Println("O usuário é maior de idade.")
	}

	var usuario2 usuario = usuario{"João", 30} 
	usuario2.apresentar() 
}
