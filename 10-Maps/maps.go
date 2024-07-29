package main

//maps: Um map é útil quando você precisa de uma coleção de pares chave-valor onde as chaves
//não são conhecidas antecipadamente ou podem variar. Maps são flexíveis, permitindo a alteração dos valores das chaves
import "fmt"

func main() {
	fmt.Println("Maps:")
	deck1 := map[string]string{ //sintaxe: var:=map[tipo]tipo {chave:valor,}
		"Picachu":     "raio",
		"Bulbasaur":   "água",
		"Chamrmander": "fogo",
	}
	fmt.Println("Picachu: Elemento:", deck1["Picachu"], "\nBulbasaur: Elemento:", deck1["Bulbasaur"])

	//maps tbm aceitam e retornam diferentes tipos comparáveis tbm
	//Usa-se map quando você precisa de flexibilidade com chaves dinâmicas ou quando não
	//conhece antecipadamente os nomes das chaves.
	//usa-se struct quando os pares chave-valor são fixos e conhecidos

	//É possível aninhar um map dentro de outro:
	usuario1 := map[string]map[string]string{
		"Nome": /*primeiro map*/ {
			"Nome":      "Juliano",
			"Sobrenome": "Silva",
		},
	}
	fmt.Println(usuario1)
	//caso seja necessário deletar alguma chave e seu valor, pois elas vêm em pares, usa-se delete(varName,"chave")
	delete(deck1, "Charmander")
	////caso seja necessário acrescentar alguma chave e seu valor, pois elas vêm em pares, usa-se
	//varName["CidadeNatal"]=map[string]string {nova chave:valor,}
	usuario1["Cidade"] = map[string]string{
		"Cidade": "Dois Irmãos",
	}
	fmt.Println(usuario1)
}
