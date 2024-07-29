package main

import "fmt"

func inverterSinal(numero int) int { // função que inverte o valor do numero fornecido
	// a função recebe um numero do tipo int e retorna um num do tipo int tbm
	return numero * -1

}

func main() {
	fmt.Println("Funções com Ponteiros")

	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
}
