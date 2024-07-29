package main

import "fmt"

// a ideia é uma função que retorna numeros na posição de fibbonaci
// 1 1 2 3 5 8 13
func fibbonaci(posicao uint) uint { // a var posição recebe apenas numeros positivos
	if posicao <= 1 {
		return posicao // aqui a condição estabelece que, se a posição for <= 1, ela retornará
		// a propria posição(0)
	}
	return fibbonaci(posicao-2) + fibbonaci(posicao-1)

}

func main() {
	fmt.Println("Recursivas") // Funções recursivas recorrem à si mesmas, ou seja, se chamam novamente
	// é necessário adicionar um ponto de parada para que a função não fique se acionando infinitamente
	//e cause um stackoverflow
	posição := uint(11)
	fmt.Println(fibbonaci(posição))
} //normalmente não se usa em numeros grande por causa do stackoverflow, pois fica empilhando muito
