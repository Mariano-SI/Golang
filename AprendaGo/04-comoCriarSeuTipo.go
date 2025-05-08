package main

import "fmt"

type hotdog int

var b hotdog =10

func main() {
	x := 20
	fmt.Printf("b: valor: %v tipo %T", b , b) //b: valor: 10 tipo main.hotdog

	b = x //erro cannot use x (variable of type int) as hotdog value in assignment
}
