package main

import "fmt"

type pessoa struct { // struct criado com seus devidos campos para informar os valores posteriormente
	nome      string
	sobrenome string
	cpf       uint64
}

type estudante struct { // nesse struct, posso chamar os campos do outro struct sem precisar digitar tudo novamente
	pessoa      // mais nada, apenas o tipo sem os campos
	numRegistro int
	curso       string
}

//OBS.: Structs podem ser criados e não usados, pois são um tipo, não uma var
func main() {
	fmt.Println("Herança")
	student1 := pessoa{"Juliano", "Silva", 0012017040}
	fmt.Println(student1)
	student1Data := estudante{student1, 0025, "Letras"} // para usar a herança, inicializo a var com o tipo(estudante) contendo o tipo (pessoa)e os valores
	fmt.Println(student1Data.curso, student1Data.numRegistro, student1Data.curso)
}
