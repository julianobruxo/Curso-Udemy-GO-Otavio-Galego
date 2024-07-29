package main

import "fmt"

func somar(n1 int, n2 int) int { // entre os () vão os parâmetros, não os valores. Após o (), declara-se o que se espera de retorno
	return n1 + n2 // valores são inseridos ao chamar a função, lá na main
}

func concat(name1, name2 string) string {
	return name1 + name2
}

func calculosMatematicos(n1, n2 int) (int, int) { // nesse caso, o retorno está entre () pq vai retornar + de 1 elemento
	soma := n1 + n2
	subtrai := n2 - n1
	return soma, subtrai
}

func main() {
	soma := somar(10, 20)
	concatenar := concat("Juliano", " da Silva")
	resultadoSoma, resultadosubtracao := calculosMatematicos(50, 25) //são 2 variaveis, uma para cada operação
	fmt.Println(soma)                                                // a var soma recebeu agora a função somar, dentro dos ()vão os valores a serem executados na função
	fmt.Println(concatenar)                                          // a var soma recebeu agora a função concat, dentro dos ()vão os valores a serem executados na função
	fmt.Println(resultadoSoma, resultadosubtracao)
	//caso eu queira usar apenas um retorno, usamos _ ao lado da var, não nos valores
}
