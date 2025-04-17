package main

import "fmt"

func main() {
	canal := make(chan string, 2) ///especifica a capacidade do canal

	canal <- "Olá, Mundo!"
	canal <- "Olá, Mundo2!"
	canal <- "Olá, Mundo2!" //volta o deadlock pois estourou a capacidade

	mensagem := <-canal

	fmt.Println(mensagem)// sem o buffer deu deadlock
}