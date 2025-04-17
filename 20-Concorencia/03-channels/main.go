package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Primeira chamada", canal)

	for{
		mensagem, aberto := <-canal //recebendo um valor
		if !aberto{
			break
		}
		fmt.Println(mensagem)
	}//gera deadlock = voce nao tem nada enviando dado pro canal mas ele ainda esta esperando. temos que tomar cuidado com deadlocks pois o go nao consegue lhes identificar em tempo de compilaçao, apenas execucao

	//sintaxe alternativa

	for mensagem := range canal{
		fmt.Println(mensagem)
	}

	fmt.Println("fim do programa")

}

func escrever(text string, channel chan string) {
	for range 5 {
		channel <- text //ao inves de printar fiz o canal receber o valor
		time.Sleep(time.Second)
	}
	//para evitar o deadlock 
	close(channel)
}