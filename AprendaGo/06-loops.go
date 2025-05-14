package main

import (
	"fmt"

)

func main() {
	//loops

	//for classico
	for x := 0; x <= 10; x++ {
		fmt.Println("Mariano", x)
	}

	//for nested loops
	for hora := 0; hora < 24; hora++{
		for minuto := 0; minuto <60; minuto++{
			for segundo := 0; segundo < 60; segundo++{
				
				fmt.Printf("%v:%v:%v\n", hora, minuto, segundo)
			}
		}
	}

	//for como while
	count := 0
	for count < 10{
		fmt.Println(count)
		count++
	}

	//for como while true

	/* 
		for{
			
		} 
	*/

	//break e continue
	for i := 0; i < 20 ; i++{
		if i % 2 != 0{
			continue //continue = quebra a iteracao atual e vai pra proxima sem paraf o loop
		}
		fmt.Println(i)
	}

	for i := 0; i < 20 ; i++{
		if i == 10{
			break //break = para o loop
		}
		fmt.Println(i)
	}
}