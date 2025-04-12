package main

import (
	"fmt"
	"time"
)

func main() {
	//concorrencia é diferente de paralelismo
	//paralelismo qcontece quando voce tem duas ou mais tarefas sendo executas ao mesmo tempo (isso so é possivel se voce tiver processador com mais de um nucleo, assim cada nucleo executa uma tarefa diferente)
	//concorrencia as tarefas nao necessariamente estao sendo executadas ao mesmo tempo, mas podem no caso de maquinas multicores, mas mesmo que voce tenha apenas um nucleo a concorencia ainda é possivel

	//como aplicar concorrencia em go
	go escrever("Olá mundo")//escrevera ola mund a cada 1 segundo Sem concorrencia ele nunca iria pra funcao de baixo vist que a execução dessa nunca acaba
	escrever("Programando em GO")//escrevera programando em go a cada 1 segundo

	//para solucionar isso podemos usar goroutines, que são funções que podem ser executadas concorrentemente com outras funções

	//quando usamos uma goroutine basicamente dizemos ao codigo para executar essa funcao em paralelo com o restante do codigo, ou seja, o programa continua executando mesmo que a goroutine ainda esteja rodando

}

func escrever(text string) {
	for{
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}