package main
//go mod init modulo
import(
	"modulo/auxiliar" 
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.go")
	auxiliar.Escrever()
	//auxiliar.escrever2() // Não compila, pois a função escrever2 não é exportada

	erro := checkmail.ValidateFormat("marianosilvagmail.com")

	fmt.Println(erro)
}