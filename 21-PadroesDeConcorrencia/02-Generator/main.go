package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo")

	for range 10{
		fmt.Println(<-canal)
	}
}

func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}