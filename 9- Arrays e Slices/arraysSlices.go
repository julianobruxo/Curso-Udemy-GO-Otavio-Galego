package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int // Arrays nada mais são do que uma lista INFLEXÍVEL de itens que, uma vez inserido o tamanho e o tipo, não podem ser mudados
	// arrays ficam salvos em uma var do tipo array + [] e dentro o tamanho do array seguido do tipo
	fmt.Println(array1) //como não foi passado nenhum valor para o array, todos terão o mesmo valor 0
	// após abrir o array sem declarar valores, pode adicionar eles usando o esquema abaixo:
	array1[2] = 5
	fmt.Println(array1)
	array1[0] = 1
	fmt.Println(array1)
	// ou abrimos {} com os valores separados por ,
	array2 := [5]int{0, 1, 2, 3, 4}
	fmt.Println(array2)
	// tbm é possível inicializar um array sem tamanho fixo nos [], mas ele vai ser limitado à quantidade de
	// valores colocados dentro das {}:
	array3 := [...]int{7, 8, 9, 6, 54, 3, 21}
	fmt.Println(array3)

	// SLICES
	//Slices são estruturas baseadas no array, porém são elásticas, ou seja, seu tamanho é dinâmico

	slice := []int{10, 65, 35, 45, 98, 45, 63, 78} // no slice, o tamanho não é especificado; ele aumenta conforme monha necessidade
	fmt.Println(slice)
	// podemos tbm "apendar" umm valor à um slice já criado usando var slice := append(var,valor)
	slice = append(slice, 15) // vai adicionar o valor 15 e retornar um novo slice
	fmt.Println(slice)
	// já que o slice é uma fatia de um array, podemos pegar, em um novo slice, apenas um pedaço do array:
	slice2 := array3[0:4] // slice2 conterá apenas os valores do indice 0 a 4 do array3
	fmt.Println(slice2)
	slice3 := array3[3:7]
	fmt.Println(slice3)
	slice4 := slice2[0:2] // tbm posso pegar valores de outro slice
	fmt.Println(slice4)
}
