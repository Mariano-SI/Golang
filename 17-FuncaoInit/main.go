package main

import "fmt"

func init() {
	fmt.Println("Executando a função init")
}

func main() {
	// Função init é uma função especial em Go que é executada antes da função main.
	// Ela é usada para inicializar pacotes ou executar código de configuração antes que o programa comece a rodar.
	// A função init não recebe parâmetros e não retorna valores.
	// Você pode ter várias funções init em um pacote, e elas serão executadas na ordem em que aparecem no código.

	fmt.Println("Executando a função main")

}