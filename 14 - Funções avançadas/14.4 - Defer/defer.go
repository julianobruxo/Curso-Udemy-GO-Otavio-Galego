package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}
func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 int) string {
	// defer fmt.Print("Aluno está aprovado: ")

	fmt.Println("Entrado na função para ver se o aluno está aprovado.")
	media := (n1 + n2) / 2
	aprovado := fmt.Sprintf("Parabéns, você passou, seu nerd de merda. Sua nota foi: %d", media)
	reprovado := fmt.Sprintf("Se fudeu, trouxa, você rodou kkkkkk. Sua nota foi: %d", media)
	if media >= 6 {
		return aprovado
	}
	return reprovado

}

func main() {
	fmt.Println("Cláusula defer")
	funcao1() // defer = adiar--> serve para adiar a execução da  funçãoselecionadasempre para o final
	funcao2()
	fmt.Println(alunoAprovado(8, 5))

}
