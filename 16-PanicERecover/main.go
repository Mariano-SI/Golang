package main

import "fmt"

func recuperarExecucao() {
	if r:= recover(); r != nil {
		fmt.Println("Recuperando a execução do programa")
	}
}

func alunoAprovado(nota1, nota2 float64) bool {
	defer recuperarExecucao() //executa antes do panic, mesmo que o panic ocorra
	media := (nota1 + nota2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6")
}

func main() {
	//panic interrompe o fluxo do programa e gera um erro, o programa para de executar e imprime o erro na tela
	//o recover serve para capturar o erro e continuar a execução do programa, ele deve ser utilizado dentro de uma função anônima (função sem nome) que é chamada com go, ou seja, em uma goroutine. OBS: antes do panic executar ele chama os defer

	fmt.Println(alunoAprovado(6,6))
	fmt.Println("Pós execução") // se a média anteriror for 6 isso nao sera executado pois a execução do programa para


}
