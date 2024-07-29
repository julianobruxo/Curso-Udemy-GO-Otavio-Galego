package main

import "fmt"

// os ponteiros servem para salvar DENTRO DE UMA VAR, não um valor, mas sim um endereço de memória para que vc possa alterar esse valor sem alterar o valor original
func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2) // var2 recebeu o valor de var 1

	variavel1++                       // dessa forma, apenas a var1 terá +1 incrementado ao seu valor, a var 2 não
	fmt.Println(variavel1, variavel2) // esse processo é atribuição de valores por CÓPIA

	// PONTEIROS SÃO UMA REFERÊNCIA DE MEMÓRIA, ou seja, ao criar uma var, vc está ocupando um
	//endereço na memória do pc e alocando a variável lá com o valor atribuído a ela
	// o PONTEIRO NÃO cria um endereço novo, ele apenas APONTA para aquele endereço de memória já existente
	//sem falar o valor.
	var variavel3 int //uma var guarda um valor inteiro
	// o valor de uma var vazia é zero
	var ponteiro *int // um ponteiro guarda um endereço de memória de um valor inteiro, referenciado pelo *
	// o valor de um ponteiro vazio é nil
	// para jogar um ponteiro dentro de uma
	// Então, como jogar um endereço de memória (ponteiro) dentro de uma var onde esse endereço foi salvo?
	// usamos & antes do nome da var
	fmt.Println(variavel3, ponteiro)
	variavel3 = 100
	ponteiro = &variavel3            // o & não vai me mostrar o valor da var3, apenas o endereço de memória que ela está
	fmt.Println(variavel3, ponteiro) // saída: valor da var 3 (var3)e o endereço de memória onde está a var 3(ponteiro)
	variavel2 = 200                  // atribuí um novo valor à var2
	ponteiro = &variavel2
	fmt.Println(variavel2, ponteiro)  // novo ponteiro tem um novo endereço agora
	fmt.Println(variavel2, *ponteiro) // ao usar *, eu desreferencio ele, dizendo que não quero o endereço mas sim o valor
}
