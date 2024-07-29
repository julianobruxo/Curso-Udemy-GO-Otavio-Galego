package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Deck Pokemon")
	deck1 := map[string]map[string]int{ // deck contém um map (pikachu) que retorna outro map(raio e poder)
		"Pikachu": {
			"Raio": 950,
		},
		"Bulbasauro": {
			"Água": 850,
		},
		"Charmander": {
			"Fogo": 900,
		},
		"Onyx": {
			"Pedra": 1050,
		},
		"Psyduck": {
			"Psíquico": 630,
		},
		"Machoke": {
			"Físico": 800,
		},
	}
	for {
		fmt.Println("Digite o nome do Pokemon que deseja pesquisar elementos e nível de poder:\nPokemons Disponíveis:\nPikachu\nBulbasauro\nCharmander\nOnyx\nPsyduck\nMachoke")
		time.Sleep(1 * time.Second)
		fmt.Println("Digite sua busca abaixo")
		time.Sleep(1 * time.Second)
		var nomePokemon string
		fmt.Scanln(&nomePokemon)

		if elementMap, exists := deck1[nomePokemon]; exists { // 'elementMap' recebe o valor associado à chave 'nomePokemon' em 'deck1'.
			// 'exists' é um booleano que indica se 'nomePokemon' está presente em 'deck1'.
			// Se 'exists' for true, esta linha imprime que o Pokémon foi encontrado,
			// mostrando o nome do Pokémon fornecido pelo usuário.

			fmt.Printf("Pokemon encontrado:%s\n", nomePokemon) // Este valor é outro mapa que mapeia elementos para seus níveis de poder (map[string]int).
			for element, power := range elementMap {           // Este loop itera sobre o 'elementMap', que recebeu o valor associado à nomePokemon
				time.Sleep(1 * time.Second)
				fmt.Printf("Elemento:%s\nNível de Poder:%d", element, power)
				// 'element' é a chave string (nome do elemento, como "Raio", "Água", etc.).
				// 'power' é o valor int (nível de poder associado ao elemento).
			}
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("Pokemon não encontrado, realize a busca novamente ")
		}
		time.Sleep(1 * time.Second)
		fmt.Print("\nDeseja buscar por outro Pokémon? (sim/não): ")
		var resposta string
		fmt.Scanln(&resposta)
		if resposta != "sim" && resposta != "Sim" {
			fmt.Println("Obrigado por usar o Deck Pokémon!")
			time.Sleep(2 * time.Second) // Adiciona um atraso de 2 segundos antes de reiniciar a busca
			break
		}

	}

}
