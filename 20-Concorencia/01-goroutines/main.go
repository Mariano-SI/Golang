package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Primeira chamada")
	escrever("Segunda chamada")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}