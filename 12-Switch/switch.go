package main

import "fmt"

func diaDaSemana(numero int) string { // a função em questão receberá um valor do tipo int e retornará uma string
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Nada"
	}
}

// tbm é possível passar mais comparações dentro do switch, mas respeitando o valor e o retorno

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1: // o controle é feito dentro do próprio switch
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Não encontrado"
	}
}

//além disso, temos a função fallthrough, que omite os demais cases indo direto para o alvo

func fimDeSemana(numero int) string {
	var diaDaSemana string

	switch {

	case numero == 1:
		diaDaSemana = "Domingo" // nesse formato, o return virá ao término da função
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
		fallthrough
	default:
		diaDaSemana = "Inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(8))
	dia := diaDaSemana(3)
	fmt.Println(dia)
}
