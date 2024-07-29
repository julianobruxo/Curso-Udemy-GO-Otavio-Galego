package main

import "fmt"

func main() { // ARITMÉTICOS
	soma := 10 + 5
	subtracao := 80 - 56
	multiply := 56 * 55
	divide := 80 / 9
	module := 20 % 3

	fmt.Println(soma, subtracao, multiply, divide, module) // todos os int devem ser do mesmo tipo, ou haverá erro(int16 + int8 = erro)

	// RELACIONAIS: Sempre retornam um bool

	fmt.Println(1 < 2) // true
	fmt.Println(2 < 1) // false
	fmt.Println(2 == 1)
	fmt.Println(2 <= 1) //menor ou igual
	fmt.Println(2 >= 1) // maior ou igual
	fmt.Println(2 != 1) // diferente

	// OPERADORES LÓGICOS
	// && --> se valor 1 E valor 2 são iguais
	// || --> se valor 1 OU valor 2 são iguais
	// ! --> diferente, inverte os valores

	// OPERADOE UNÁRIO
	numero := 1
	numero++     // vai uncrementar o valor em 1
	numero += 2  // aqui vai incrementar(+) no numero igual ao valor (=15)
	numero--     // vai DEcrementar o valor em 1
	numero -= 15 // aqui vai DEncrementar(+) no numero igual ao valor (=15)
	numero *= 3  // aqui vai pegar o numero e fazer numero * 3

	fmt.Println(numero)

	if numero > 5 {
		fmt.Println("Número é maior que 5")
	} else {
		fmt.Println("Número é menor que 5")
	}
	//OPERADORE TERNÁRIO --> não existe em GO; usa-se if / else

	//texto:= numero >5 ? "Maior que 5" : "Menor que 5"

}
