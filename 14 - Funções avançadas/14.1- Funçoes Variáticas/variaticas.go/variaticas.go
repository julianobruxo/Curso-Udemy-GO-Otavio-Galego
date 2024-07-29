package main

import "fmt"

// funções variáticas servem para poder passar vários números ao mesmo tempo para a mesma função
// para isso, não especificamos a quantidade de números nos parametros, mas sim var ... tipo
// Ex.: func calculo(numeros...int){}
// posso passar outros parametros comos strings ou structs tbm
// Deve ser ir por último

func soma(numeros ...int) int { // esse tipo de função retorna um slice
	//sendo um slice, podemos iterar sobre ele e somar todos os numeros dentro:
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
func main() {
	fmt.Println("Funçoes variáticas")
	totalDaSoma := soma(1, 5) // pode passar quantos numeros quiser somar ou nenhum
	fmt.Println(totalDaSoma)
}
