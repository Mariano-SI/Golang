package main

import "fmt"

type retangulo struct {
	altura  float64
	largura float64
}

// area implements forma.
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

// area implements forma.
func (c circulo) area() float64 {
	return c.raio * c.raio * 3.14
}

type forma interface { //interface é um tipo que define um conjunto de metodos que devem ser implementados por um tipo
	area() float64
}

//sem interface se eu quisesse uma funcao que escrevesse a area de cada forma teria que na verdade ter duas funções, uma pra cada.
// agora como defini uma interface que tem o método area, posso passar qualquer tipo que implemente essa interface e a função vai funcionar
// ou seja, a interface é um tipo que define um contrato, ou seja, define quais métodos um tipo deve ter para ser considerado daquele tipo.
// é interesasnte que em momento algum fizemos uma ligacao explicita entre as structs e a interface, diferente de outras lags que geralmente usam a clausula implements

func escreverArea(f forma) {
	fmt.Println("A area é: ", f.area())
}
func main() {
	r := retangulo{altura: 10, largura: 5}
	escreverArea(r) // A area é:  50

	c := circulo{raio: 10}	
	escreverArea(c) // A area é:  314
}
