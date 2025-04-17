package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for{
			time.Sleep(time.Millisecond * 500)//meio segundo
			canal1 <- "Canal1"
		}
	}()
	go func() {
		for{
			time.Sleep(time.Second * 2) //dois segundos
			canal2 <- "Canal2"
		}
	}()

	for{
		//solucao = select
		//;select evitar forçar uma operacao concorrente depender de outra pois isso tecnicamente atrapalharia a concorrencia

		select {
		case mensagemCanal1 := <- canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <- canal2:
			fmt.Println(mensagemCanal2)
		}

		//isso gera perca de tempo pois cada mensagem do canal1 espera uma do canal2 sendo que enquanto esperamos uma mensagem do canal 2 dava pra rodar 4 mensagens do 1

		/* 
		mensagemCanal1 := <- canal1
		fmt.Println(mensagemCanal1)

		mensagemCanal2 := <- canal2
		fmt.Println(mensagemCanal2)
		
		*/
	}
}