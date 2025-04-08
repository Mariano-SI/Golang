package main

func calculosMatematicos(a int, b int) (soma int, subtracao int){
	soma = a + b
	subtracao = a - b
	return //retorna os valores na ordem que foram declarados
}

func main() {
	//funcoes com retorno nomeado

	soma, subtracao := calculosMatematicos(10, 5)
	println(soma, subtracao) // 15 5
}