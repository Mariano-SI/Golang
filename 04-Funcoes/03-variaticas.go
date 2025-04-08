package main 

func somaVariatica(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func escrever(text string, numbers ...int){
	for _, number := range numbers {
		println(text, number)
	}
}

func main(){
	//funcoes variaticas sao funcoes que podem receber um numero variavel de argumentos
	//funcao variatica recebe o ... antes do tipo de dado, e dentro da funcao podemos usar como um slice normal

	resultado := somaVariatica(10, 2, 3, 4, 5)
	println(resultado) // 24

	escrever("Ola", 1, 2, 3, 4, 5)
}