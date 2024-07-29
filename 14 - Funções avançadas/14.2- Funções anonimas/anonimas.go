package main

import "fmt"

//funções anonimas são executadas e encerradas logo após. São funções que não precisam nome,
//pois não são chamadas fora do escopo

func main() {
	func() {
		fmt.Println("Anonymous func") //posso passar o parametro aqui ou
	}()
	func(texto string) { //passar o parametro aqui
		fmt.Println(texto)
	}("Anonymous 2") //ou aqui
}
