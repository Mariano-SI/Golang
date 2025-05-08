package main

import "fmt"

type hotdog int

var b hotdog =10

func main() {
	x := 20
	fmt.Printf("x: valor: %v tipo %T\n", x , x) //x: valor: 20 tipo int
	fmt.Printf("b: valor: %v tipo %T\n", b , b) //b: valor: 10 tipo main.hotdog

	//x = b //cannot use b (variable of type hotdog) as int value in assignment

	x = int(b) //converti  hotdog pra int

	fmt.Printf("x: valor: %v tipo %T\n", x , x)//x: valor: 10 tipo int


}
