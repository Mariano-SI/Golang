package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //quantidade de routines

	go func(){
		escrever("Primeira chamada")
		waitGroup.Done() //-1
	}()
	go func(){
		escrever("Segunda chamada")
		waitGroup.Done()//1
	}()

	waitGroup.Wait() //so vai encerrar a main se todas as rotinas estiverem finalizadas
}

func escrever(text string) {
	for range 5{
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}