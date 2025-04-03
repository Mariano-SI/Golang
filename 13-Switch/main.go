package main

import "fmt"

//primeira forma de usar switch
func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "sabado"
	default:
		return "Número inválido"
	}
}

//segunda maneira

func diaDaSemana2(num int) string{
	switch{
	case num == 1:
		return "Domingo"
	case num == 2:
		return "Segunda"
	case num == 3:
		return "Terça"
	case num == 4:
		return "Quarta"
	case num == 5:
		return "Quinta"
	case num == 6:
		return "Sexta"
	case num == 7:
		return "sabado"
	default:
		return "Número inválido"
	}
}

func main() {
	dia := diaDaSemana(1)
	dia2 := diaDaSemana2(4)

	fmt.Println(dia)
	fmt.Println(dia2)

	//falar da clausula fallthrough
}