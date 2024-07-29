package main // Structs são uma coleção de campos, onde cada um rtem um nome e um tipo

import "fmt"

// type --> identifica o tipo do campo que eu preciso, como usuário, documento, data de nasc., etc
type usuario struct { // usuario é o nome do tipo que estou criando, sendo essa a sintaxe básica
	// aqui dentro vou criar os tipos que eu quero e declarar o que cada um é,ou seja, o tipo do valor, se é string ou int ou o que for
	nome          string // o campo nome será do tipo string
	idade         int    // o campo idade será do tipo int
	nacionalidade string
	telefone      int
}

type endereco struct {
	rua    string
	numero int
}

func main() {
	fmt.Println("Arquivo structs")
	// como criar variaveis do tipo usuario?
	var user usuario // agora usuário, que eu criei, é um tipo, assim como string ou int ou bool(agora subtipos)
	// para atribuir daos aos campos, usamos o nome da  var.campo = "Valor"
	user.nome = "Juliano\n"
	user.idade = 40
	user.nacionalidade = "Brasileira\n"
	user.telefone = 51999309238
	fmt.Println(user)
	// tbm é possível fornecer os dados para o struct diretamente após declarar a var dentro de {}--> var := nomeCampo{dado1,"dado2", dao3...}
	user2 := usuario{"Isabelle", 4, "Brasileira", 51999309239}
	fmt.Println(user2)

	// tbm é possível passar um valor de outro struct para dentro do primeiro struct
	adress := endereco{"Avenida São Miguel", 1210} //--> lembre-se que o tipo criado pelo struct deve ser posto dntro de uma var antes de ser usado
	user3 := usuario{"Suzane", 40, "Brasileira", 51997585792}
	fmt.Println(user3, adress)
}
