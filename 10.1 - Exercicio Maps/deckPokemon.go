package main

import "fmt"

func main() {
	fmt.Println("Deck Pockemon")
	deck1 := map[string]int{
		"Raio": 950,
		"Água": 850,
		"Fogo": 900,
	} // na saída, é preferível usar PrintF e placeholders para buscar os valors, e listar os maps ao final
	// como boa prática
	fmt.Printf("Níveis de Poder:\n Nível de Poder Picachu (Raio): %d, \n Nível de Poder: Bulbasaur (Água):%d,\n Nível de Poder do Charmander(Fogo):%d", deck1["Raio"], deck1["Água"], deck1["Fogo"])
	// para adicionar um item ao map desaninhado, usamos varName["nome da chave"] = valor da chave
	deck1["Pedra"] = 1050
	fmt.Printf("Níveis de Poder:\n Nível de Poder Picachu (Raio): %d, \n Nível de Poder: Bulbasaur (Água):%d,\n Nível de Poder do Charmander(Fogo):%d,\n Nível de Poder do Onyx(Pedra): %d", deck1["Raio"], deck1["Água"], deck1["Fogo"], deck1["Pedra"])
}
