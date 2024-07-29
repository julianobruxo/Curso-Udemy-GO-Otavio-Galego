package main

import (
	"fmt"
)

func main() {
	// i := 0       // A variável i é inicializada com 0
	// for i < 10 { // Enquanto i for menor que 10
	// 	i += 2                         // Incrementa i em 2 a cada iteração
	// 	fmt.Println("Incrementando i") // Imprime uma mensagem indicando que i foi incrementado
	// 	time.Sleep(1 * time.Second)    // Espera 1 segundo
	// }
	// fmt.Println(i) // Imprime o valor final de i

	// for j := 10; j < 15; j++ { // Inicializa j com 10; enquanto j for menor que 15; incrementa j em 1 a cada iteração
	// 	fmt.Println("Metendo j", j) // Imprime uma mensagem junto com o valor atual de j
	// 	time.Sleep(1 * time.Second) // Espera 1 segundo
	// 	//fmt.Println(j)//erro, pois ela não está visível no escopo por ser declarada de forma curta
	// }

	nomes := [3]string{"João", "Carla", "Maria"}
	for indice, nome := range nomes { // aqui ídice é o index do valor, nomes são as strings
		// se eu não quisere retornar o índice, faço for_,nomes:= range no campo que quero percorrer
		fmt.Printf("O índice é %d e o valor é %s\n", indice, nome)
	}
	comMaps := map[string]string{ // criei o map
		"nome":      "Marcia",
		"sobrenome": "Silva",
	}
	for _, valor := range comMaps { // darei um RANGE e listarei a chave e o valor da chave
		fmt.Println(valor)
	}
}
