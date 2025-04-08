package main

func main(){
	//funcoes anonimas sao funcoes que nao tem nome e podem ser atribuida a variaveis, passadas como argumento ou retornadas de outras funcoes
	func(){
		println("Olá, Mundo!")
	}() //chamada da funcao anonima


	soma := func(x, y int) int {
		return x + y
	}
	println(soma(10, 20)) // 30
	//funcao anonima atribuida a variavel e chamada
}