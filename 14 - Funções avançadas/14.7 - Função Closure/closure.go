package main

import "fmt"

// a função closure é uma função que referencia variáveis q estão fora do seu corpo
//como funções são tipos, podemos retornar elas como qualquer outro tipo

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() { //a var funcao recebeu uma função anonima sem retorno que printa a var texto
		fmt.Println(texto)
	}
	return funcao // preciso retornar a funcao para que ela possa ser usada fora do escopo da func closure
}
func main() {
	texto := "Dentro da Função main" // essa variavel texto NÃO É a mesma da func closure
	fmt.Println(texto)
	funcNova := closure() //criar a var funcNova que recebe a func closure, a qual printa "Dentro da função closure"
	//a var funcNova virou uma função ao receber a func closure:
	funcNova() // closure fez com que a var funcNova se tornasse uma função, portanto ela recebe as atribuições da func
	//isso é util para trazer uma função de outro escopo para dentro de um novo escopo,
	//já que funcs são restritas ao seus escopos originais
}
