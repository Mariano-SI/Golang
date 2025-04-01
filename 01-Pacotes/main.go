package main
//go mod init modulo
import(
	"modulo/auxiliar" 
	"fmt"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.go")
	auxiliar.Escrever()
	//auxiliar.escrever2() // Não compila, pois a função escrever2 não é exportada
}