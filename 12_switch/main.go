package main

import "fmt"

func diaSemana(num int) string{
	switch num{
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terca"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero invalido"
	}
}

func test(num int) string{
	switch{
		case num == 1:
			return "Test"
		default:
			return"Test"
	}
}

func main() {
	fmt.Println("Switch")

	dia:= diaSemana(4)
	fmt.Println(dia)
}