package main

import "fmt"

func funcao1() {
	println("Executando a funcao1")
}
func funcao2() {
	println("Executando a funcao2")
}

func alunoAprovado(nota1, nota2 float32)bool{
	defer fmt.Println(("Média calculada, o resultado sera retornado"))//executa antes do return da funcao
	fmt.Println("Entrando na função para ver se o aluno está aprovado")
	media := (nota1 + nota2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	//a clausula serve para adiar a execução de um pedaço de codigo ate o ultimo momento dentro de um escopo
	//se o defer for aplicado em uma funcao com retorno, o defer executa imediatamente antes do retorno

	defer funcao1() // 3- por ultimo essa
	funcao2() //1- primeiro executa essa

	fmt.Println("Executando a funcao main") //2 - depois essa

	fmt.Println(alunoAprovado(7,8))

	//defer é muito utilizado para fechar conexões, arquivos, etc. POr exemplo para fechar conexoes com banco de dados. Evita a repetição da sentença de fechamento em varios lugares
}