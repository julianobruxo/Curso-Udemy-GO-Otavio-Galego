package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	// essa função irá receber 2 parametros: n1 e n2 do tipo int e retornar soma e sub dos 2 n
	soma = n1 + n2 // pq não foi inicializado com := ? pq já está no retorno, não precisa inicializar
	sub = n1 - n2
	return // vai retornar o que está entre () (soma int, sub int)
	// normalmente apenas o tip é colocado entre os () e as {}, mas como são 2 operações, vai um novo ()
	// contendo as operações
}

func main() {
	fmt.Println("Retorno nomeado")
	fmt.Println(calculosMatematicos(25, 30))
}
