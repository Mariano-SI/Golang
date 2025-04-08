package main

import "errors"

func somar(a int8, b int8) int8 {
	return a + b
}
func subtrair(a int8, b int8) int8 {
	return a - b
}
func multiplicar(a int8, b int8) int8 {
	return a * b
}
func dividir(a int8, b int8) (int8, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

func CalculosMatematicos(a int8, b int8) (int8, int8, int8, int8, error) {
	soma := somar(a, b)
	subtracao := subtrair(a, b)
	multiplicacao := multiplicar(a, b)
	divisao, err := dividir(a, b)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return soma, subtracao, multiplicacao, divisao, nil
}

func main() {
	soma := somar(10, 5)
	subtracao := subtrair(10, 5)
	multiplicacao := multiplicar(10, 5)
	divisao, err := dividir(10, 0)
	if err != nil {
		println(err.Error())
	} else {
		println(divisao)
	}
	println(soma)
	println(subtracao)
	println(multiplicacao)

	var f = func(){
		println("Funcao anonima")
	}

	f()
}