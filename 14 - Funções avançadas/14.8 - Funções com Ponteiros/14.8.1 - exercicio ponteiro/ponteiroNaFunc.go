package main

import "fmt"

func desenhar(cor *string) {
	*cor = "azul" // ponteiro que aponta para a var cor, não está inicializada, está apenas recebendo o valor azul
	fmt.Println("Desenhando com a cor", *cor)
}

func main() {
	cor := "vermelho"
	desenhar(&cor) // Passa um ponteiro para 'cor' para a função
	fmt.Println("A cor original após a função: ", cor)
}
