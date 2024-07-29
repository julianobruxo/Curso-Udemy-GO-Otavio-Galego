package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
} // dessa forma, podemos passar qualquer tipo para a interface, não apenas metodos

func main() {
	fmt.Println("Interfaces genéricas")
	generica("String")
	generica(50)
	generica(true)
	generica(2.5)

}
