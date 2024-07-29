package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	idade     int
}

// sintaxe: func (nome da var(u) nome do tipo criado no struct(usuario)) o nome do método (salvar()) {}
func (u usuario) salvar() { // um método é uma função que retorna outra função e trabalha com
	//structs e interfaces
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)

}

func (u usuario) maiorDeIdade() bool { // esse metodo verifica se o dado "idade" dentro do struct é true ou false
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Juliano", "Silva", 41}
	fmt.Println(usuario1)
	usuario1.salvar() // aqui eu chamei o método, que apenas imprime a string
	// poderia ser qualquer função associada à ususario 1
	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario1.aniversario()
	fmt.Println(usuario1.idade)
}
