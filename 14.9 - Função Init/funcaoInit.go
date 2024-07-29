package main

import "fmt"

func main() {
	fmt.Println("Rodando a main")
}

func init() {
	fmt.Println("Rodando a init") // func init sempre será executada antes das demais,
	//e não importa onde ela esteja
}
