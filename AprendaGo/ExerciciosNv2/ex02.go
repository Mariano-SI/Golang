package main

import "fmt"

func main() {
	a := 1 == 2
	b := 3 <= 8
	c := 45 >= 8
	d := 12 < 11
	e := 33 > 1
	f := 5 != 5

	fmt.Println("1 == 2", a)
	fmt.Println("3 <= 8", b)
	fmt.Println("45 >= 8", c)
	fmt.Println("12 < 11", d)
	fmt.Println("33 > 1", e)
	fmt.Println("5 != 5", f)
}