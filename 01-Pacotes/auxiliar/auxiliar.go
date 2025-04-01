package auxiliar

import "fmt"

// Funcoes acessiveis para outros pacotes devem começar com letra maiúscula
// Funcoes que não devem ser acessadas fora do pacote devem começar com letra minuscula
func Escrever(){
	fmt.Println("Escrevendo do arquivo auxiliar.go")
	escrever2() // Chama a função escrever2, que é acessível dentro do mesmo pacote
}