package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle em Go:")
	// primeiro cria-se a var que levara o valor a ser comparado se é true ou false:
	numero := 15

	if numero > 15 {
		fmt.Println("É maior que 15")
	} else {
		fmt.Println("É menor ou igual a 15")
	}
	// if init = quando a estrutura de controle é usada apenas para uma variável simples, ela não pode ser
	//usada fora do escopo do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("É maior que 0")
	} else {
		fmt.Println("É menor ou igual a 0")
	}
	//fmt.Println(outroNumero)// erro, pois está fora do escopo do if init
}
