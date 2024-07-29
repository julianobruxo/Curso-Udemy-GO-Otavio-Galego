package main

import "fmt"

func main() {
	fmt.Println("Arrays internos")
	// arrays internos
	// primeiro será criado um make;
	interno1 := make([]float32, 10, 11) //var = make(tipo(slice[])]tipo de dado, tamanho, tamanho máximo)
	fmt.Println(interno1)
	fmt.Println(len(interno1)) // len mostra o tamanho do array/slice/map sendo usado
	fmt.Println(cap(interno1)) //len mostra a capacidade do array/slice/map sendo usado
	interno1 = append(interno1, 3)
	fmt.Println(interno1) //aumentamos o número de itens dentro do slice, atingindo a capacidade máxima
	fmt.Println(len(interno1))
	fmt.Println(cap(interno1))
	interno1 = append(interno1, 5) //ao apendar mais, o comprimento e a capacidade máxima do slice dobram
	fmt.Println(len(interno1))
	fmt.Println(cap(interno1))

	interno2 := make([]int, 5) // make sem passar a capacidade máxima: O go deixa a capacidade máxima
	fmt.Println(interno2)      // igual ao tamanho que vc passou
	fmt.Println(len(interno2))
	fmt.Println(cap(interno2))
	interno2 = append(interno2, 1) //ao apendar mais um item, o slice passou para 6 intens e a capacidade dobrou para 10
	fmt.Println(interno2)
	fmt.Println(len(interno2))
	fmt.Println(cap(interno2))

}
